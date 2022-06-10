.PHONY: test build proto kind

race:
	mkdir -p artifacts
	go test -race -short -cover -coverprofile=artifacts/coverage.txt -covermode=atomic ./...

docker:
	docker build -t ghcr.io/jack-hughes/users-service:local-dev .

up:
	docker-compose up --build -d
	curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ \
	-d @scripts/debezium/register-postgres.json

down:
	docker-compose down -v --remove-orphans

proto:
	protoc --go_out="pkg" --go-grpc_out="pkg" \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
		api/users/users.proto

generate: proto
	go mod tidy
	go mod vendor
	go generate ./...
	go fmt ./...
	go vet ./...

lint:
	golangci-lint run ./... --timeout=5m

test:
	go test ./...

helm-template:
	helm template users-service charts/

kind-up: docker
	./scripts/kind.sh

kind-down:
	helm delete local-release -n postgres
	kubectl delete pvc -l release=local-release
	kind delete cluster --name users-service

build-cli:
	go build -o bin/userctl ./cmd/userctl

integration: up build-cli
	./scripts/integration.sh