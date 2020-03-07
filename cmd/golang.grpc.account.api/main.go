package main

import (
	"github.com/gorilla/mux"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/seongsukang/golang.grpc.account.api/app/interface/rpc/v1.0/protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

const (
	serverHost = "127.0.0.1"
	serverPort = ":7777"
	certFile   = "./cert/my_private.key.crt"
	KeyFile    = "./cert/my_private.key"
)

type AccountServer struct{}

func (as AccountServer) Create(context.Context, *account.CreateRequest) (*account.CreateResponse, error) {
	panic("implement me")
}

func (as AccountServer) Read(context.Context, *account.ReadRequest) (*account.ReadResponse, error) {
	panic("implement me")
}

func (as AccountServer) Update(context.Context, *account.UpdateRequest) (*account.UpdateResponse, error) {
	panic("implement me")
}

func (as AccountServer) Delete(context.Context, *account.DeleteRequest) (*account.DeleteResponse, error) {
	panic("implement me")
}

func recoveryHandler(p interface{}) (err error) {
	return status.Errorf(codes.Unknown, "panic triggered: %v", p)
}

func main() {
	serverCert, err := credentials.NewServerTLSFromFile(certFile, KeyFile)
	if err != nil {
		panic(err)
	}

	clientCert, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		panic(err)
	}

	recOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoveryHandler),
	}

	opts := []grpc.ServerOption{}
	opts = append(opts, grpc.Creds(serverCert))
	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_recovery.UnaryServerInterceptor(recOpts...),
	))

	grpcServer := grpc.NewServer(opts...)

	account.RegisterAccountServiceServer(grpcServer, new(AccountServer))

	restRouter := runtime.NewServeMux()
	conn, err := grpc.DialContext(
		context.Background(),
		serverHost,
		grpc.WithTransportCredentials(clientCert),
	)
	if err != nil {
		panic(err)
	}

	if err := account.RegisterAccountServiceHandler(context.Background(), restRouter, conn); err != nil {
		panic(err)
	}

	httpHandler := mux.NewRouter()
	httpHandler.PathPrefix("/account").Handler(httpGrpcRouter(grpcServer, httpHandler))

	http.ListenAndServeTLS(serverPort, certFile, KeyFile, httpHandler)
}

func httpGrpcRouter(grpcServer http.Handler, httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	})
}
