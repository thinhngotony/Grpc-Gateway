.PHONY: install, clean, swagger, run, auto

# protoc --go_out=plugins=grpc:. -I proto/ proto/*.proto --grpc-gateway_out=logtostderr=true:.
install:
# Make latest gRPC code
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:.
	protoc --proto_path=proto proto/*.proto --grpc-gateway_out=logtostderr=true:.

clean:
	rm -rf proto/protogen/*.go

swagger:
	protoc --proto_path=proto proto/*.proto --swagger_out=proto/swagger

build:
	sh build.sh

run:
	make clean && make install && make build && ./grpc-gateway-server

# --swagger_out=logtostderr=true:${OUT_PATH}