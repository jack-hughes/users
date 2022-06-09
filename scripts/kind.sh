#!/bin/bash

# Bring up the kind cluster
kind create cluster --name users-service

# Load the users service image
mkdir -p $HOME/tmp/
TMPDIR=$HOME/tmp/ kind load docker-image ghcr.io/jack-hughes/users-service:local-dev --name users-service

# Create init.sql
kubectl create namespace postgres
kubectl create configmap db-config -n postgres --from-file=./scripts/db/init.sql

## Install Postgres
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install local-release bitnami/postgresql -n postgres \
--set auth.postgresPassword=postgres \
--set auth.database=users \
--set primary.initdb.scriptsConfigMap=db-config \
--wait

# Install the users-service dependencies (secrets obviously unsafe here)
kubectl create namespace users
kubectl create secret generic db-secret \
  --from-literal=username=postgres \
  --from-literal=password=postgres \
  --namespace users

# Install local helm chart
helm install users-service ./charts -n users --wait