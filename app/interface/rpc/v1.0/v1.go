package v1

import (
	"github.com/nodias/golang.grpc.account.api/app/interface/rpc/v1.0/account"
	"github.com/nodias/golang.grpc.account.api/app/registry"
	"github.com/nodias/golang.grpc.account.api/app/usecase"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	account.RegisterAccountServiceServer(server, NewUserService(ctn.Resolve("user-usecase").(usecase.UserUsecase)))
}
