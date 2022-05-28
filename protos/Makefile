# Go API Design With Protocol Buffers and gRPC
# https://betterprogramming.pub/go-api-design-with-protocol-buffers-and-grpc-991838e4852b

# install swagger and protoc-gen-openapiv2:
install:
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

# make proto annotation_file=posts_annotations.yml proto_file=posts.proto
annotation_file=posts_annotations.yml
proto_file=posts.proto

proto:
	mkdir -p src/go/sma src/swagger src/openapi
	protoc --go_out=src/go/sma --go_opt=paths=source_relative \
        --go-grpc_out=src/go/sma --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=logtostderr=true,paths=source_relative,grpc_api_configuration=$(annotation_file):src/go/sma \
        --swagger_out=grpc_api_configuration=$(annotation_file):src/swagger \
        --openapiv2_out=logtostderr=true:src/openapi $(proto_file)
