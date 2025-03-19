tidy:
	@go mod tidy

# gen-proto:
# 	@protoc --go_out=proto/__generated__ --go_opt=paths=source_relative \
#   --go-grpc_out=proto/__generated__ --go-grpc_opt=paths=source_relative \
#   proto/*.proto

gen-proto:
	@protoc --go_out=proto/__generated__ --go_opt=module=github.com/iamNilotpal/grpc/proto \
  --go-grpc_out=proto/__generated__ --go-grpc_opt=module=github.com/iamNilotpal/grpc/proto \
  proto/*.proto

run-hello-server:
	@go build -o bin/hello/server/main cmd/hello/server/main.go
	@./bin/hello/server/main

run-hello-client:
	@go build -o bin/hello/client/main cmd/hello/client/main.go
	@./bin/hello/client/main

run-todo-server:
	@go build -o bin/todo/server/main cmd/todo/server/main.go
	@./bin/todo/server/main

run-todo-client:
	@go build -o bin/todo/client/main cmd/todo/client/main.go
	@./bin/todo/client/main