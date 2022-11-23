package main

import (
	model "BairesDev/Level6/02_Final_Microservices/users/grpc/model"
	mygrpc "BairesDev/Level6/02_Final_Microservices/users/grpc/service"
	mytransport "BairesDev/Level6/02_Final_Microservices/users/rest/http"
	mylog "BairesDev/Level6/02_Final_Microservices/users/rest/log"
	myrest "BairesDev/Level6/02_Final_Microservices/users/rest/service"

	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/go-kit/log"
	grpc "google.golang.org/grpc"
)

func ServeGrpc(port string, dbType string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(fmt.Sprintf("Error Listening for gRPC on Port:%s", port))
	}

	s := mygrpc.NewGrpcUserService(dbType)

	grpcServer := grpc.NewServer()
	model.RegisterUsersServer(grpcServer, s)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", port, "caller", log.DefaultCaller)

	logger.Log("msg", "GRPC", "addr", port)
	logger.Log("err", grpcServer.Serve(listener))
}

func ServeRest(port string, dbType string) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", port, "caller", log.DefaultCaller)

	svc := mylog.NewLoggingMiddleware(logger, myrest.NewService(dbType))

	//Migrate Struct(s) to DB Tables
	svc.UserRESTService.Migrate(context.Background())

	r := mytransport.NewHttpServer(svc, logger)
	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func main() {
	//Db Type (mysql, redis)
	dbType := "mysql"

	//Serve Users Via REST Server On Port 8001
	go ServeRest("8001", dbType)

	//Serve Users Via GRPC Server On Port 9001
	go ServeGrpc("9001", dbType)

	for true {
	}
}
