FROM golang:1.18.2-alpine3.15 AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /users-service  /build/cmd/server/

RUN GRPC_HEALTH_PROBE_VERSION=v0.3.2 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

FROM alpine:3.15.4
RUN apk --no-cache add ca-certificates
COPY --from=builder /users-service /opt
COPY --from=builder /bin/grpc_health_probe /grpc_health_probe
EXPOSE 5355
CMD ["/opt/users-service"]
