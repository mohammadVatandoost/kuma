FROM gcr.io/distroless/static-debian11:debug-nonroot
# FROM golang:1.19 as base_image


# FROM ubuntu:latest

COPY /tools/releases/templates/LICENSE \
    /tools/releases/templates/README \
    /tools/releases/templates/NOTICE \
    /kuma/

# WORKDIR /build-app
COPY . .


# RUN ls
# RUN go mod download
# RUN make build
# RUN ls 
# RUN cd /test && ls
# WORKDIR /kuma
# RUN ls
# RUN pwd

SHELL ["/busybox/busybox", "sh", "-c"]
