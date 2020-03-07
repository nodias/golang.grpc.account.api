package main

import (
	"fmt"
	"github.com/gorilla/mux"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/seongsukang/golang.grpc.account.api/app/interface/rpc/v1.0/protocol"
	_ "github.com/seongsukang/golang.grpc.account.api/statik"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

const (
	serverHost = "127.0.0.1"
	serverPort = ":443"
	MongoHost  = "127.0.0.1"
	MongoPort  = ":27017"
	certFile   = "./cert/my_private.key.crt"
	KeyFile    = "./cert/my_private.key"
)

type AccountServer struct{}

func (a AccountServer) CreateUser(context.Context, *account.CreateUserRequest) (*account.CreateUserResponse, error) {
	panic("implement me")
}

func (a AccountServer) ReadUsers(*account.ReadUsersRequest, account.AccountService_ReadUsersServer) error {
	panic("implement me")
}

func (a AccountServer) ReadUser(context.Context, *account.ReadUserRequest) (*account.ReadUserResponse, error) {
	panic("implement me")
}

func (a AccountServer) UpdateUser(context.Context, *account.UpdateUserRequest) (*account.UpdateUserResponse, error) {
	panic("implement me")
}

func (a AccountServer) DeleteUser(context.Context, *account.DeleteUserRequest) (*account.DeleteUserResponse, error) {
	panic("implement me")
}

func recoveryHandler(p interface{}) (err error) {
	return status.Errorf(codes.Unknown, "panic triggered: %v", p)
}

var url = fmt.Sprintf("%s%s", serverHost, serverPort)

func main() {
	//db, err := InitMongoDB(MongoHost, MongoPort)
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

	var opts []grpc.ServerOption
	opts = append(opts, grpc.Creds(serverCert))
	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_recovery.UnaryServerInterceptor(recOpts...),
	))

	grpcServer := grpc.NewServer(opts...)

	account.RegisterAccountServiceServer(grpcServer, new(AccountServer))

	conn, err := grpc.DialContext(
		context.Background(),
		url,
		grpc.WithTransportCredentials(clientCert),
	)
	if err != nil {
		panic(err)
	}

	restRouter := runtime.NewServeMux()
	if err := account.RegisterAccountServiceHandler(context.Background(), restRouter, conn); err != nil {
		panic(err)
	}

	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	fs := http.StripPrefix("/swaggerui", http.FileServer(statikFS))

	mux := mux.NewRouter()
	mux.PathPrefix("/swaggerui").Handler(fs)
	mux.PathPrefix("/account/users").Handler(restRouter)

	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		panic(err)
	}

	srv := http.Server{
		Addr: url,
		Handler: httpGrpcRouter(grpcServer, mux),
	}

	go func() {
		log.Printf("## Start server port: %s", serverPort)
		srv.ServeTLS(lis, certFile, KeyFile)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("## Stopping server...")
	srv.Close()
	log.Println("## Bye.")
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
