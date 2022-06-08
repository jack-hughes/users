docker:
	docker build -t ghcr.io/jack-hughes/users:local-dev .

up:
	docker-compose up --build

down:
	docker-compose down -v

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