package postgres

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/attribute"

	config "github.com/kumahq/kuma/pkg/config/plugins/resources/postgres"
	core_model "github.com/kumahq/kuma/pkg/core/resources/model"
	"github.com/kumahq/kuma/pkg/core/resources/store"
	core_metrics "github.com/kumahq/kuma/pkg/metrics"
	pgx_config "github.com/kumahq/kuma/pkg/plugins/resources/postgres/config"
)

type pgxResourceStore struct {
	pool *pgxpool.Pool
}

var _ store.ResourceStore = &pgxResourceStore{}

// This attribute is necessary for tracing integrations like Datadog, to have
// full insights into sql queries connected with traces.
// ref. https://github.com/DataDog/dd-trace-go/blob/3d97fcec9f8b21fdd821af526d27d4335b26da66/contrib/database/sql/conn.go#L290
var spanTypeSQLAttribute = attribute.String("span.type", "sql")

func NewPgxStore(metrics core_metrics.Metrics, config config.PostgresStoreConfig, customizer pgx_config.PgxConfigCustomization) (store.ResourceStore, error) {
	pool, err := connect(config, customizer)
	if err != nil {
		return nil, err
	}

	if err := registerMetrics(metrics, pool); err != nil {
		return nil, errors.Wrapf(err, "could not register DB metrics")
	}

	return &pgxResourceStore{
		pool: pool,
	}, nil
}

func connect(postgresStoreConfig config.PostgresStoreConfig, customizer pgx_config.PgxConfigCustomization) (*pgxpool.Pool, error) {
	connectionString, err := postgresStoreConfig.ConnectionString()
	if err != nil {
		return nil, err
	}
	pgxConfig, err := pgxpool.ParseConfig(connectionString)

	if postgresStoreConfig.MaxOpenConnections == 0 {
		// pgx MaxCons must be > 0, see https://github.com/jackc/puddle/blob/c5402ce53663d3c6481ea83c2912c339aeb94adc/pool.go#L160
		// so unlimited is just max int
		pgxConfig.MaxConns = math.MaxInt32
	} else {
		pgxConfig.MaxConns = int32(postgresStoreConfig.MaxOpenConnections)
	}
	pgxConfig.MinConns = int32(postgresStoreConfig.MinOpenConnections)
	pgxConfig.MaxConnIdleTime = time.Duration(postgresStoreConfig.ConnectionTimeout) * time.Second
	pgxConfig.MaxConnLifetime = postgresStoreConfig.MaxConnectionLifetime.Duration
	pgxConfig.MaxConnLifetimeJitter = postgresStoreConfig.MaxConnectionLifetime.Duration
	pgxConfig.HealthCheckPeriod = postgresStoreConfig.HealthCheckInterval.Duration
	pgxConfig.ConnConfig.Tracer = otelpgx.NewTracer(otelpgx.WithAttributes(spanTypeSQLAttribute))
	customizer.Customize(pgxConfig)

	if err != nil {
		return nil, err
	}

	return pgxpool.NewWithConfig(context.Background(), pgxConfig)
}

func (r *pgxResourceStore) Create(ctx context.Context, resource core_model.Resource, fs ...store.CreateOptionsFunc) error {
	opts := store.NewCreateOptions(fs...)

	bytes, err := core_model.ToJSON(resource.GetSpec())
	if err != nil {
		return errors.Wrap(err, "failed to convert spec to json")
	}

	var ownerName *string
	var ownerMesh *string
	var ownerType *string

	if opts.Owner != nil {
		ptr := func(s string) *string { return &s }
		ownerName = ptr(opts.Owner.GetMeta().GetName())
		ownerMesh = ptr(opts.Owner.GetMeta().GetMesh())
		ownerType = ptr(string(opts.Owner.Descriptor().Name))
	}

	version := 0
	statement := `INSERT INTO resources VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`
	_, err = r.pool.Exec(ctx, statement, opts.Name, opts.Mesh, resource.Descriptor().Name, version, string(bytes),
		opts.CreationTime.UTC(), opts.CreationTime.UTC(), ownerName, ownerMesh, ownerType)
	if err != nil {
		if strings.Contains(err.Error(), duplicateKeyErrorMsg) {
			return store.ErrorResourceAlreadyExists(resource.Descriptor().Name, opts.Name, opts.Mesh)
		}
		return errors.Wrapf(err, "failed to execute query: %s", statement)
	}

	resource.SetMeta(&resourceMetaObject{
		Name:             opts.Name,
		Mesh:             opts.Mesh,
		Version:          strconv.Itoa(version),
		CreationTime:     opts.CreationTime,
		ModificationTime: opts.CreationTime,
	})
	return nil
}

func (r *pgxResourceStore) Update(ctx context.Context, resource core_model.Resource, fs ...store.UpdateOptionsFunc) error {
	bytes, err := core_model.ToJSON(resource.GetSpec())
	if err != nil {
		return err
	}

	opts := store.NewUpdateOptions(fs...)

	version, err := strconv.Atoi(resource.GetMeta().GetVersion())
	newVersion := version + 1
	if err != nil {
		return errors.Wrap(err, "failed to convert meta version to int")
	}
	statement := `UPDATE resources SET spec=$1, version=$2, modification_time=$3 WHERE name=$4 AND mesh=$5 AND type=$6 AND version=$7;`
	result, err := r.pool.Exec(
		ctx,
		statement,
		string(bytes),
		newVersion,
		opts.ModificationTime.UTC(),
		resource.GetMeta().GetName(),
		resource.GetMeta().GetMesh(),
		resource.Descriptor().Name,
		version,
	)
	if err != nil {
		return errors.Wrapf(err, "failed to execute query %s", statement)
	}
	if rows := result.RowsAffected(); rows != 1 {
		return store.ErrorResourceConflict(resource.Descriptor().Name, resource.GetMeta().GetName(), resource.GetMeta().GetMesh())
	}

	// update resource's meta with new version
	resource.SetMeta(&resourceMetaObject{
		Name:             resource.GetMeta().GetName(),
		Mesh:             resource.GetMeta().GetMesh(),
		Version:          strconv.Itoa(newVersion),
		ModificationTime: opts.ModificationTime,
	})

	return nil
}

func (r *pgxResourceStore) Delete(ctx context.Context, resource core_model.Resource, fs ...store.DeleteOptionsFunc) error {
	opts := store.NewDeleteOptions(fs...)

	statement := `DELETE FROM resources WHERE name=$1 AND type=$2 AND mesh=$3`
	result, err := r.pool.Exec(ctx, statement, opts.Name, resource.Descriptor().Name, opts.Mesh)
	if err != nil {
		return errors.Wrapf(err, "failed to execute query: %s", statement)
	}
	if rows := result.RowsAffected(); rows == 0 {
		return store.ErrorResourceNotFound(resource.Descriptor().Name, opts.Name, opts.Mesh)
	}

	return nil
}

func (r *pgxResourceStore) Get(ctx context.Context, resource core_model.Resource, fs ...store.GetOptionsFunc) error {
	opts := store.NewGetOptions(fs...)

	statement := `SELECT spec, version, creation_time, modification_time FROM resources WHERE name=$1 AND mesh=$2 AND type=$3;`
	row := r.pool.QueryRow(ctx, statement, opts.Name, opts.Mesh, resource.Descriptor().Name)

	var spec string
	var version int
	var creationTime, modificationTime time.Time
	err := row.Scan(&spec, &version, &creationTime, &modificationTime)
	if err == pgx.ErrNoRows {
		return store.ErrorResourceNotFound(resource.Descriptor().Name, opts.Name, opts.Mesh)
	}
	if err != nil {
		return errors.Wrapf(err, "failed to execute query: %s", statement)
	}

	if err := core_model.FromJSON([]byte(spec), resource.GetSpec()); err != nil {
		return errors.Wrap(err, "failed to convert json to spec")
	}

	meta := &resourceMetaObject{
		Name:             opts.Name,
		Mesh:             opts.Mesh,
		Version:          strconv.Itoa(version),
		CreationTime:     creationTime.Local(),
		ModificationTime: modificationTime.Local(),
	}
	resource.SetMeta(meta)

	if opts.Version != "" && resource.GetMeta().GetVersion() != opts.Version {
		return store.ErrorResourcePreconditionFailed(resource.Descriptor().Name, opts.Name, opts.Mesh)
	}
	return nil
}

func (r *pgxResourceStore) List(ctx context.Context, resources core_model.ResourceList, args ...store.ListOptionsFunc) error {
	opts := store.NewListOptions(args...)

	statement := `SELECT name, mesh, spec, version, creation_time, modification_time FROM resources WHERE type=$1`
	var statementArgs []interface{}
	statementArgs = append(statementArgs, resources.GetItemType())
	argsIndex := 1
	if opts.Mesh != "" {
		argsIndex++
		statement += fmt.Sprintf(" AND mesh=$%d", argsIndex)
		statementArgs = append(statementArgs, opts.Mesh)
	}
	if opts.NameContains != "" {
		argsIndex++
		statement += fmt.Sprintf(" AND name LIKE $%d", argsIndex)
		statementArgs = append(statementArgs, "%"+opts.NameContains+"%")
	}
	statement += " ORDER BY name, mesh"

	rows, err := r.pool.Query(ctx, statement, statementArgs...)
	if err != nil {
		return errors.Wrapf(err, "failed to execute query: %s", statement)
	}
	defer rows.Close()

	total := 0
	for rows.Next() {
		item, err := rowToItem(resources, rows)
		if err != nil {
			return err
		}
		if err := resources.AddItem(item); err != nil {
			return err
		}
		total++
	}

	resources.GetPagination().SetTotal(uint32(total))
	return nil
}

func rowToItem(resources core_model.ResourceList, rows pgx.Rows) (core_model.Resource, error) {
	var name, mesh, spec string
	var version int
	var creationTime, modificationTime time.Time
	if err := rows.Scan(&name, &mesh, &spec, &version, &creationTime, &modificationTime); err != nil {
		return nil, errors.Wrap(err, "failed to retrieve elements from query")
	}

	item := resources.NewItem()
	if err := core_model.FromJSON([]byte(spec), item.GetSpec()); err != nil {
		return nil, errors.Wrap(err, "failed to convert json to spec")
	}

	meta := &resourceMetaObject{
		Name:             name,
		Mesh:             mesh,
		Version:          strconv.Itoa(version),
		CreationTime:     creationTime.Local(),
		ModificationTime: modificationTime.Local(),
	}
	item.SetMeta(meta)

	return item, nil
}

func (r *pgxResourceStore) Close() error {
	r.pool.Close()
	return nil
}

func registerMetrics(metrics core_metrics.Metrics, pool *pgxpool.Pool) error {
	postgresCurrentConnectionMetric := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "store_postgres_connections",
		Help: "Current number of postgres store connections",
		ConstLabels: map[string]string{
			"type": "open_connections",
		},
	}, func() float64 {
		return float64(pool.Stat().TotalConns())
	})

	postgresInUseConnectionMetric := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "store_postgres_connections",
		Help: "Current number of postgres store connections",
		ConstLabels: map[string]string{
			"type": "in_use",
		},
	}, func() float64 {
		return float64(pool.Stat().AcquiredConns())
	})

	postgresIdleConnectionMetric := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "store_postgres_connections",
		Help: "Current number of postgres store connections",
		ConstLabels: map[string]string{
			"type": "idle",
		},
	}, func() float64 {
		return float64(pool.Stat().IdleConns())
	})

	postgresMaxOpenConnectionMetric := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "store_postgres_connections_max",
		Help: "Max postgres store open connections",
	}, func() float64 {
		return float64(pool.Stat().MaxConns())
	})

	postgresWaitConnectionMetric := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "store_postgres_connection_wait_count",
		Help: "Current waiting postgres store connections",
	}, func() float64 {
		return float64(pool.Stat().EmptyAcquireCount())
	})

	postgresWaitConnectionDurationMetric := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "store_postgres_connection_wait_duration",
		Help: "Time Blocked waiting for new connection in seconds",
	}, func() float64 {
		return pool.Stat().AcquireDuration().Seconds()
	})

	postgresMaxIdleClosedConnectionMetric := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "store_postgres_connection_closed",
		Help: "Current number of closed postgres store connections",
		ConstLabels: map[string]string{
			"type": "max_idle_conns",
		},
	}, func() float64 {
		return float64(pool.Stat().MaxIdleDestroyCount())
	})

	postgresMaxLifeTimeClosedConnectionMetric := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "store_postgres_connection_closed",
		Help: "Current number of closed postgres store connections",
		ConstLabels: map[string]string{
			"type": "conn_max_life_time",
		},
	}, func() float64 {
		return float64(pool.Stat().MaxLifetimeDestroyCount())
	})

	postgresSuccessfulAcquireCountMetric := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "store_postgres_connection_acquire",
		Help: "Cumulative count of acquires from the pool",
		ConstLabels: map[string]string{
			"type": "successful",
		},
	}, func() float64 {
		return float64(pool.Stat().AcquireCount())
	})

	postgresCanceledAcquireCountMetric := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "store_postgres_connection_acquire",
		Help: "Cumulative count of acquires from the pool",
		ConstLabels: map[string]string{
			"type": "canceled",
		},
	}, func() float64 {
		return float64(pool.Stat().CanceledAcquireCount())
	})

	postgresConstructingConnectionsCountMetric := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "store_postgres_connections",
		Help: "Current number of postgres store connections",
		ConstLabels: map[string]string{
			"type": "constructing",
		},
	}, func() float64 {
		return float64(pool.Stat().ConstructingConns())
	})

	postgresNewConnectionsCountMetric := prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "store_postgres_connections",
		Help: "Current number of postgres store connections",
		ConstLabels: map[string]string{
			"type": "new",
		},
	}, func() float64 {
		return float64(pool.Stat().NewConnsCount())
	})

	if err := metrics.
		BulkRegister(postgresCurrentConnectionMetric, postgresInUseConnectionMetric, postgresIdleConnectionMetric,
			postgresMaxOpenConnectionMetric, postgresWaitConnectionMetric, postgresWaitConnectionDurationMetric,
			postgresMaxIdleClosedConnectionMetric, postgresMaxLifeTimeClosedConnectionMetric, postgresSuccessfulAcquireCountMetric,
			postgresCanceledAcquireCountMetric, postgresConstructingConnectionsCountMetric, postgresNewConnectionsCountMetric,
		); err != nil {
		return err
	}
	return nil
}
