.PHONY: gen dev

gen:
	@protoc -I=proto --go_out=module=authzed-go:. \
	--go-grpc_out=module=authzed-go:. \
	./proto/*.proto


dev:
	@go run main.go