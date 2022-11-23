package main

import (
	model "BairesDev/Level6/02_Final_Microservices/answers/grpc/model"
	mygrpc "BairesDev/Level6/02_Final_Microservices/answers/grpc/service"
	mytransport "BairesDev/Level6/02_Final_Microservices/answers/rest/http"
	mylog "BairesDev/Level6/02_Final_Microservices/answers/rest/log"
	myrest "BairesDev/Level6/02_Final_Microservices/answers/rest/service"

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

	s := mygrpc.NewGrpcAnswerService(dbType)

	grpcServer := grpc.NewServer()
	model.RegisterAnswersServer(grpcServer, s)

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
	svc.AnswerRESTService.Migrate(context.Background())

	r := mytransport.NewHttpServer(svc, logger)
	logger.Log("msg", "HTTP", "addr", port)
	logger.Log("err", http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func main() {
	//Db Type (mysql, redis)
	dbType := "mysql"

	//Serve Users Via REST Server On Port 8001
	go ServeRest("8003", dbType)

	//Serve Users Via GRPC Server On Port 9001
	go ServeGrpc("9003", dbType)

	for true {
	}
}
