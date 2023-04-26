import{M as V}from"./kongponents.es-5ca9e130.js";import{d as q,u as C,r,v as L,y as M,A as P,j as b,i as h,g as k,w as z,b as x,h as A,f as N,B as D,o as d,e as F}from"./index-6ef061d4.js";import{_ as Q}from"./ZoneEgressDetails.vue_vue_type_script_setup_true_lang-7b692b45.js";import{D as R}from"./DataOverview-273e56b2.js";import{u as U}from"./index-8eaa4fe5.js";import{Q as Z}from"./QueryParameter-70743f73.js";import"./AccordionList-fc691476.js";import"./_plugin-vue_export-helper-c27b6911.js";import"./DefinitionListItem-97bb646e.js";import"./ErrorBlock-5ce00c3e.js";import"./LoadingBlock.vue_vue_type_script_setup_true_lang-431ef500.js";import"./EnvoyData-a689fa08.js";import"./CodeBlock.vue_vue_type_style_index_0_lang-cadc6770.js";import"./StatusInfo.vue_vue_type_script_setup_true_lang-dfb8a1dd.js";import"./SubscriptionHeader.vue_vue_type_script_setup_true_lang-60a557a2.js";import"./TabsWidget-a1cf8c62.js";import"./datadogLogEvents-302eea7b.js";import"./TagList-2830199d.js";import"./store-444aa12f.js";import"./StatusBadge-310603a4.js";const $={class:"zoneegresses"},j={class:"kcard-stack"},G={class:"kcard-border"},K={key:0,class:"kcard-border"},fe=q({__name:"ZoneEgresses",props:{selectedZoneEgressName:{type:String,required:!1,default:null},offset:{type:Number,required:!1,default:0}},setup(T){const i=T,f=U(),B={title:"No Data",message:"There are no Zone Egresses present."},u=C(),n=r(!0),o=r(!1),c=r(null),m=r({headers:[{label:"Status",key:"status"},{label:"Name",key:"name"}],data:[]}),l=r(null),g=r([]),y=r(null),O=r([]),_=r(i.offset);L(()=>u.params.mesh,function(){u.name==="zone-egress-list-view"&&(n.value=!0,o.value=!1,c.value=null,p(0))}),M(function(){p(i.offset)});async function p(t){_.value=t,Z.set("offset",t>0?t:null),n.value=!0,o.value=!1;const s=u.query.ns||null,a=D;try{const{data:e,next:v}=await I(s,a,t);y.value=v,e.length?(o.value=!1,g.value=e,E({name:i.selectedZoneEgressName??e[0].name}),m.value.data=e.map(w=>{const S=P(w.zoneEgressInsight??{});return{...w,status:S}})):(m.value.data=[],o.value=!0)}catch(e){e instanceof Error?c.value=e:console.error(e),o.value=!0}finally{n.value=!1}}function E({name:t}){var a;const s=g.value.find(e=>e.name===t);if(s){const e=((a=s.zoneEgressInsight)==null?void 0:a.subscriptions)??[];O.value=Array.from(e).reverse(),l.value=s,Z.set("zoneEgress",t)}}async function I(t,s,a){if(t)return{data:[await f.getZoneEgressOverview({name:t},{size:s,offset:a})],next:null};{const{items:e,next:v}=await f.getAllZoneEgressOverviews({size:s,offset:a});return{data:e??[],next:v}}}return(t,s)=>{var a;return d(),b("div",$,[h("div",j,[h("div",G,[k(R,{"selected-entity-name":(a=l.value)==null?void 0:a.name,"page-size":x(D),"is-loading":n.value,error:c.value,"empty-state":B,"table-data":m.value,"table-data-is-empty":o.value,next:y.value,"page-offset":_.value,onTableAction:E,onLoadData:p},{additionalControls:z(()=>[t.$route.query.ns?(d(),F(x(V),{key:0,class:"back-button",appearance:"primary",icon:"arrowLeft",to:{name:"zone-egress-list-view"}},{default:z(()=>[A(`
              View all
            `)]),_:1})):N("",!0)]),_:1},8,["selected-entity-name","page-size","is-loading","error","table-data","table-data-is-empty","next","page-offset"])]),A(),l.value!==null?(d(),b("div",K,[k(Q,{"zone-egress-overview":l.value},null,8,["zone-egress-overview"])])):N("",!0)])])}}});export{fe as default};
