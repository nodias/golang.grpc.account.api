PROJECT_PATH:=/Users/seongsukang/Documents/develop/workspace/golang-module/TOY-grpc-kafka-simpleChatApp/back-end/golang.grpc.account.api
GOPATH:=/Users/seongsukang/Documents/develop/workspace/golang-workspace/gopath
GRPC_ECOSYSTEM:=${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.13.0
GRPC_ECOSYSTEM_GATEWAY:=${GRPC_ECOSYSTEM}/protoc-gen-grpc-gateway
GRPC_ECOSYSTEM_SWAGGER:=${GRPC_ECOSYSTEM}/protoc-gen-swagger
GRPC_ECOSYSTEM_GATEWAY_PROTOPATH:=${GRPC_ECOSYSTEM}/third_party/googleapis
GRPC_ECOSYSTEM_SWAGGER_PROTOPATH:=${GRPC_ECOSYSTEM}
MODULE_STATIK:=${GOPATH}/pkg/mod/github.com/rakyll/statik@v0.1.7

## This makefile is gRPC, REST transcoding basic setup guide + Running guide of my TOY project
## by seongsukang(nodias46@gmail.com, github.com/nodias)

all:


.PHONY:module-setup
module-setup: module-setup-grpc-ecosystem-gateway module-setup-grpc-ecosystem-swagger module-setup-filesystem-statik module-setup-filesystem-statik-init utility-swagger-setup
	@echo MAKE : module-setup
	go list -m all

module-setup-grpc-ecosystem-gateway:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
	cd ${GRPC_ECOSYSTEM_GATEWAY} && \
 	go install

module-setup-grpc-ecosystem-swagger:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && \
	cd ${GRPC_ECOSYSTEM_SWAGGER} && \
 	go install

module-setup-filesystem-statik:
	go get -u github.com/rakyll/statik && \
	cd ${MODULE_STATIK} && \
	go install && \
	go list -m all && \

module-setup-filesystem-statik-init:
	cp -rp api/swagger swaggerui/apis && \
	${GOPATH}/bin/statik -src=${PROJECT_PATH}/swaggerui

utility-swagger-setup:
	@echo MAKE : utility-swagger-setup
	git clone https://github.com/swagger-api/swagger-ui.git && \
	mv swagger-ui/dist ./swaggerui && \
	rm -rf swagger-ui
	#ls | grep swagger-ui | awk '{print "rm -rf " $0}'|sh -v

.PHONY:proto-generate
proto-generate:
	@echo MAKE : proto-generate
	protoc \
	--proto_path=api/proto/v1.0 \
	--proto_path=${GRPC_ECOSYSTEM_GATEWAY_PROTOPATH} \
	--proto_path=${GRPC_ECOSYSTEM_SWAGGER_PROTOPATH} \
	--go_out=plugins=grpc:app/interface/rpc/v1.0/account \
	--grpc-gateway_out=logtostderr=true:app/interface/rpc/v1.0/account \
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




