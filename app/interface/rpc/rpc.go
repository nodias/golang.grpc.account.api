package rpc

import (
	"github.com/nodias/golang.grpc.account.api/app/interface/rpc/v1.0"
	"github.com/nodias/golang.grpc.account.api/app/registry"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	v1.Apply(server, ctn)
}
