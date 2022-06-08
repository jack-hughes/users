FROM golang:1.18.2-alpine3.15 AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /users-service  /build/cmd/server/

FROM alpine:3.15.4
RUN apk --no-cache add ca-certificates
COPY --from=builder /users-service /opt
EXPOSE 5355
CMD ["/opt/users-service"]
