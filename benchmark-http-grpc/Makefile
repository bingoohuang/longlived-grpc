
bench:
	go test -bench=. -benchmem=1  -benchtime=30s

pb:
	protoc -I./grpc --go_out=. --go-grpc_out=. users.proto

init:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
