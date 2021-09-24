package main

import (
	"exaple.com/gRPC/grpctest"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct{
	grpctest.UnimplementedGrpctestServer
}

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	grpctest.RegisterGrpctestServer(grpcServer,&Server{})
	if err := grpcServer.Serve(lis); err!=nil{
		log.Fatalf("Failed to server port 9000 %v",err)
	}

}