generate-auth-api:
	protoc --proto_path api/auth_v1 \
	--go_out=pkg/auth_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go.exe \
	--go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc.exe \
	api/auth_v1/auth.proto

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.13.0

proto-gen:
	docker run -it --rm -v "//D:/DevGO//hh-auth/://app" -w "//app" rvolosatovs/protoc protoc \
		--proto_path=./api/auth_v1 \
		--go_out=pkg/auth_v1 \
		--go_opt=paths=source_relative \
		--go-grpc_out=pkg/auth_v1 \
		--go-grpc_opt=paths=source_relative \
		auth.proto