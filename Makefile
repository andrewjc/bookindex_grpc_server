install-protobuf:
	brew install protobuf

install-protoc-gen-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

install-protoc-gen-go-grpc:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

install: install-protobuf install-protoc-gen-go install-protoc-gen-go-grpc set-go-path proto

proto:
	protoc --go_out=./internal/grpc --go_opt=paths=source_relative \
        --go-grpc_out=./internal/grpc --go-grpc_opt=paths=source_relative \
        proto/*.proto

set-go-path:
	export PATH="$PATH:$(go env GOPATH)/bin"

run:
	go run cmd/main.go

querytest:
	grpcurl -plaintext -d '{"query": "book"}' localhost:8080 BookIndex/Search
