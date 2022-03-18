protoc:
	protoc --go_out=plugins=grpc:. protos/*.proto
build:
	go install ./...
