#!/bin/bash

mkdir -p artifacts
set -e

docker-compose up --build -d

# Get the container ip for the service
ip=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' connect)
echo "container ip: $ip"

curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://$ip:8083/connectors/ \
-d @scripts/debezium/register-postgres.json