.PHONY: all

proto:
	protoc --proto_path=api/v1/proto --go_out=api/v1/proto/gen --go_opt=paths=source_relative \
	--go-grpc_out=api/v1/proto/gen --go-grpc_opt=paths=source_relative api/v1/proto/*.proto