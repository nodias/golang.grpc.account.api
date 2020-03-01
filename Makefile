GOPATH:=/Users/seongsukang/Documents/develop/workspace/golang-workspace/gopath
GRPC_ECOSYSTEM:=${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.13.0
GRPC_ECOSYSTEM_GATEWAY:=${GRPC_ECOSYSTEM}/protoc-gen-grpc-gateway
GRPC_ECOSYSTEM_SWAGGER:=${GRPC_ECOSYSTEM}/protoc-gen-swagger
GRPC_ECOSYSTEM_GATEWAY_PROTOPATH:=${GRPC_ECOSYSTEM}/third_party/googleapis
GRPC_ECOSYSTEM_SWAGGER_PROTOPATH:=${GRPC_ECOSYSTEM}

all: module-setup proto-generate

.PHONY:module-setup
module-setup:
	@echo MAKE : module-setup
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && \
    go list -m all && \
	cd ${GRPC_ECOSYSTEM_GATEWAY} && \
 	go install && \
	cd ${GRPC_ECOSYSTEM_SWAGGER} && \
 	go install

.PHONY:proto-generate
proto-generate:
	@echo MAKE : proto-generate
	protoc \
	--proto_path=api/proto/v1.0 \
	--proto_path=${GRPC_ECOSYSTEM_GATEWAY_PROTOPATH} \
	--proto_path=${GRPC_ECOSYSTEM_SWAGGER_PROTOPATH} \
	--go_out=plugins=grpc:app/interface/rpc/v1.0/protocol \
	--grpc-gateway_out=logtostderr=true:app/interface/rpc/v1.0/protocol \
	--swagger_out=logtostderr=true:api/swagger/v1.0 \
	account.proto

.PHONY:cert-generate cert-privateKey cert-rm_pp cert-csr cert-selfsigned
cert-generate-selfsigned: cert-privateKey cert-rm_pp cert-csr cert-selfsigned
	@echo MAKE : cert-generate cert-rm-pp

cert-privateKey:
	@echo MAKE : pass_phrase \for example key is '1234'
	openssl genrsa -des3 -out ./cert/my_private.key 2048

cert-rm_pp:
	@echo MAKE : remove pass_phrase
	cp ./cert/my_private.key ./cert/my_private.key.enc
	openssl rsa -in ./cert/my_private.key.enc -out ./cert/my_private.key
	chmod 600 ./cert/my_private.key*

cert-csr:
	@echo MAKE : csr
	openssl req -new -key ./cert/my_private.key -out ./cert/my_private.key.csr

cert-selfsigned:
	@echo MAKE : self-singed key
	openssl x509 -req \
	-days 3650 \
	-in ./cert/my_private.key.csr \
	-signkey ./cert/my_private.key \
	-extfile ./cert/test.com.cnf \
	-out ./cert/my_private.key.crt

cert-check-crt:
	@echo MAKE : check selfsigned key
	@echo Subject Alternative Name is iportant!!
	openssl x509 -text -in ./cert/my_private.key.crt

cert-check-csr:
	@echo MAKE : check csr
	openssl req -text -in ./cert/my_private.key.csr




