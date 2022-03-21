protoc:
	protoc --go_out=plugins=grpc:. protos/*.proto
install:
	go install ./...
