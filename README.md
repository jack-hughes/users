# users

[![codecov](https://codecov.io/gh/jack-hughes/users/branch/main/graph/badge.svg?token=8X6JNWP97U)](https://codecov.io/gh/jack-hughes/users)
[![CI](https://github.com/jack-hughes/users/actions/workflows/service.yaml/badge.svg)](https://github.com/jack-hughes/users/actions/workflows/service.yaml)

## Table of Contents

- [Overview](#overview)
    - [Features](#features)
- [Requirements](#requirements)
- [Tooling](#tooling)
    - [Makefile](#makefile)
    - [CLI](#cli)
- [Local Deployment](#local-deployment)
    - [Kubernetes](#kubernetes)
        - [Health Checks](#health-checks)
    - [docker-compose](#docker-compose)
    - [Database](#database)
    - [Images](#images)
- [Events](#events)
- [Testing](#testing)
    - [Unit Testing](#unit-testing)
    - [Integration Testing](#integration-testing)
- [Assumptions and Decisions](#assumptions-and-decisions)
- [Possible Extensions and Improvements](#possible-extensions-and-improvements)

## Overview

Users is a small gRPC service that allows CRUD operations
against a user entity. The service is implemented in Go and
uses a Postgres database for the backend data storage. It
also provides a friendly and easy to use CLI tool
called `userctl` to interact with the gRPC service.

### Features

- Full CI suite including:
    - Linting
    - Unit testing + code coverage reporting
    - Integration testing using `userctl`
    - Build and push to the GitHub container registry
- Health checking endpoints
- Helm Charts for deployment to Kubernetes
- Bespoke CLI tooling - `userctl`
- 95% unit test coverage
- Events emitted on user entity creation, deletion, update
  and reads
- docker-compose local environment

## Requirements

- [Go](https://go.dev/) (built using version 1.18.2)
- [protoc](https://grpc.io/docs/languages/go/quickstart/)
  and respective Go tooling
- [Helm](https://helm.sh/docs/intro/install/)
- [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/)
- [golangci-lint](https://golangci-lint.run/usage/install/)
- [Docker](https://docs.docker.com/engine/install/)
- [docker-compose](https://docs.docker.com/compose/install/)

## Tooling

### Makefile

At the root of the project directory, you will find a
Makefile that will enable you to run various helpful tasks
within the project. Find a table describing them below

| Command       | Functionality                                                                      |
|---------------|------------------------------------------------------------------------------------|
| race          | Run unit tests and generate code coverage                                          |
| docker        | Build and tag a docker image for the users-service                                 |
| up            | Bring up a local docker-compose environment                                        |
| down          | Bring down a local docker-compose environment and remove volumes                   |
| proto         | Generate code from proto definitions                                               |
| lint          | Run golangci against the codebase                                                  |
| helm-template | Test generating the helm deployment and service to stdout                          |
| kind-up       | Build a local docker image and install helm templates against a local kind cluster |
| kind-down     | Tear down the local kind cluster and uninstall helm charts                         |
| build-cli     | Compile the `userctl` cli tool (found in the bin directory)                        |
| integration   | Bring up the docker compose stack and run integration tests against it             |

### CLI

This project also delivers a CLI tool to interact with the
gRPC server. To build it - use the `make build-cli` tool as
described above. By executing `./bin/userctl` you will be
presented with a helpful screen for how to use the CLI tool:

```shell
A friendly CLI for interacting with the users service

Usage:
  userctl [command]

Available Commands:
  create      create a user
  delete      delete a user based on their id
  help        Help about any command
  list        lists users based on a filter of a 2 character country code
  update      update a user

Flags:
  -h, --help          help for userctl
  -a, --host string   host of the gRPC Server (default "localhost")
  -p, --port string   port of the gRPC Server (default "5355")

Use "userctl [command] --help" for more information about a command
```

## Local Deployment

The users service provides two methods of deploying the
stack locally. The first is via `Kind` using `Helm` charts (
see the `charts/` directory), with the second being
via `docker-compose`.
The two options are available because:

- Using Kind allows us to test kubernetes deployments,
  services, readiness and liveliness checks
- Using Kind allows us to test our Helm chart
  implementations

Deploying via Kind does not deploy the full stack - only the
service, Postgres (and its config including `initdb`
scripts) to provide a minimalst Kubernetes test environment.

Deploying the `docker-compose` stack locally provides a
larger more well rounded stack for more comprehensive local
development. This stack deploys:

- The Users service
- Postgres (initialising the dataset)
- Kafka (for relaying events)
- Zookeeper (tracks Kafka nodes and status)
- Kafka Connect (allows integrations between Postgres and
  Kafka)
- Kafdrop (WebUI for viewing Kafka topics and messages)
- alpine/curl (to configure Kafka Connect)

### Kubernetes

To deploy via to Kubernetes locally using kind run:
`make kind-up`. This will perform the following steps:

- Build the user service Docker image
- Bring up a Kind cluster
- Load the built service image into Kind
- Create an initdb SQL configmap for Postgres
- Install the Postgres Helm chart
- Create Postgres secrets
- Install the users-service Helm chart

Note that this process may take a while, particularly on
first time usage as it pulls images from registries.
When successfully deployed, you should be able to port
forward using kubectl to the service and be able to make
requests. You can use either `userctl` or a tool
like [Bloom](https://github.com/bloomrpc/bloomrpc) to
interact with the service.

Example:

`kubectl port-forward -n users svc/users-service 5355:5355`

Run `make kind-down` to bring down the local cluster and
cleanup Helm charts.

#### Health Checks

A probe is installed that interacts with the health check
endpoints at the point of building the Docker image.
Kubernetes then leverages this via the deployment as seen
below:

```yaml
readinessProbe:
  exec:
    command: [ "/grpc_health_probe", "-addr=:{{ .Values.env.grpcPort }}" ]
  initialDelaySeconds: 5
livenessProbe:
  exec:
    command: [ "/grpc_health_probe", "-addr=:{{ .Values.env.grpcPort }}" ]
```

### Docker Compose

To bring up the compose environment, simply run `make up`.
This will bring up the aforementioned services.
When the images have pulled and are running, I highly
recommend interacting with the service either
using `userctl` or BloomRPC whilst inspecting the Kafdrop
WebUI.

**Note**: you may need to find the IP address of the users
service to interact with it. To do this,
run: `docker inspect users-service` and grab the IP address
from the Networks.IPAddress field.

To access the Kafdrop WebUI, navigate to `localhost:9000`.
Here, you can see the available topics on the Kafka instance
running in docker-compose.
Topics are auto-created when events are pushed to them in
this local instance, so use your preferred tooling to create
a new user against the service.

```shell
./bin/userctl create -a 192.168.32.7 --first-name=test --last-name=test --nickname=test --password=123456 --email=test@devv.com --country=BU`
```

When you've created the user, refresh the Kafdrop page and
you should see a new topic called db.users.users:
![](https://i.imgur.com/fbrCRiL.png)

Click on this and hit view messages. You will be presented
with events on CRUD operations on the user entity, which can
later be consumed by other services that may care.

To bring down the `docker-compose` stack, run `make down`.

### Database

The database is managed by Postgres' inbuilt initdb scripts.
These SQL files will be run on the initial boot of Postgres
and are
configured [here](https://github.com/jack-hughes/users/blob/main/scripts/db/init.sql)
.

Things to note:

- User IDs are UUIDs and the unique primary key of the table
- Emails are unique fields
- A trigger is created to automatically update
  the `updated_at` timestamps on row changes
- The database is initially seeded with three fake users

### Images

As part of the CI process, images are built and pushed on PR
and merge to main. You can find those
images [here](https://github.com/jack-hughes/users/pkgs/container/users)
.

## Events

This section will discuss my approach to emitting events
from the users-service. When I created this project I wanted
to keep a singular microservice that would cover the domain
of modifying and creating user entities.
This proposes a problem. If in this instance the workflow
is:

**Client -> Create User -> Store in Database**

At what point can I safely emit an event with the guarentee
that both the database row will be stored and delivered to
Kafka?
I toyed with the idea of the following approach:

**Client -> Create User -> Store in Database -> Success ->
Produce message to Kafka in code**

However, if the Kafka connectivity drops after successfully
storing the row in the database, this is unrecoverable.

I opted for the approach that the database itself would emit
events. I chose this option because debezium will _always_
reconcile rows with corresponding Kafka messages, meaning
that we will never lose a CRUD operation.

Kafka Connect and Debezium have significant operational
overhead and are hard to manage. These technologies are fine
for a small, singular microservice.
If I had more time, and wanted to make a more complex
project, I would emit user creation events before storing
them in the database and completely removing the database
component from the users microservice. When a creation event
is emitted from the microservice, I would also have a
database storage consumer which would persist the event to
an RDBMS, as well as any other listening microservices that
cared for changes to a user entity.

## Testing
### Unit Testing
- All business logic is unit tested with code coverage reports to >95%
### Integration Testing
- Integration testing uses the `userctl` command line tool against the docker-compose stack to ensure all CRUD operations work as expected.

## Assumptions and Decisions
- Storage and service interfaces are designed to allow swapping out the database storage technology easily without affecting service functionality.
- Passwords are hashed when stored in the database
- Emails are assumed to be unique in the user entity

## Possible Extensions and Improvements
- Continuous delivery with Flux or ArgoCD to a staging/production Kubernetes cluster
- Migration tooling to manage database schema changes
- Certificate based authentication rather than username and password authentication for databases
- Helm charts pushed to a remote repository