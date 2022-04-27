.DEFAULT_GOAL := gen, clean

gen:
	protoc --proto_path=proto proto/*.proto --go_out=internal/service/proto
	protoc --proto_path=proto proto/*.proto --go-grpc_out=internal/service/proto

clean:
	rm internal/service/proto/*.go