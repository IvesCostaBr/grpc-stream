generate:
	@protoc --go_out=. --go-grpc_out=. ./proto/core.proto

build:
	@echo "---- Building Application ----"
	@go build -o server-bin cmd/server/*.go
	# @go build -o client-bin client/*.go

run:
	@echo "---- Running Server ----"
	@go run cmd/server/*.go

run_client:
	@echo "---- Running Client ----"
	@go run client/*.go