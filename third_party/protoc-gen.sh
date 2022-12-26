
protoc --proto_path=api/proto/v1 --go_out=. \
    --go-grpc_out=. todo-service.proto


