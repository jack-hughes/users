proto:
	protoc --go_out="pkg" --go-grpc_out="pkg" \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
		api/users/users.proto

docker:
	docker build -t ghcr.io/jack-hughes/api:local-dev .

up:
	docker-compose up --build

down:
	docker-compose down -v