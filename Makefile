GOPATH:=/Users/seongsukang/Documents/develop/workspace/golang-workspace/gopath
GRPC_ECOSYSTEM:=${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.13.0
GRPC_ECOSYSTEM_GATEWAY:=${GRPC_ECOSYSTEM}/protoc-gen-grpc-gateway
GRPC_ECOSYSTEM_SWAGGER:=${GRPC_ECOSYSTEM}/protoc-gen-swagger
GRPC_ECOSYSTEM_GATEWAY_PROTOPATH:=${GRPC_ECOSYSTEM}/third_party/googleapis
GRPC_ECOSYSTEM_SWAGGER_PROTOPATH:=${GRPC_ECOSYSTEM}

all: module-setup proto-generate

.PHONY:module-setup
module-setup:
	@echo MAKE - module-setup
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && \
    go list -m all && \
	cd ${GRPC_ECOSYSTEM_GATEWAY} && \
 	go install && \
	cd ${GRPC_ECOSYSTEM_SWAGGER} && \
 	go install

.PHONY:proto-generate
proto-generate:
	protoc \
	--proto_path=api/proto/v1.0 \
	--proto_path=${GRPC_ECOSYSTEM_GATEWAY_PROTOPATH} \
	--proto_path=${GRPC_ECOSYSTEM_SWAGGER_PROTOPATH} \
	--go_out=plugins=grpc:app/interface/rpc/v1.0/protocol \
	--grpc-gateway_out=logtostderr=true:app/interface/rpc/v1.0/protocol \
	--swagger_out=logtostderr=true:api/swagger/v1.0 \
	account.proto
