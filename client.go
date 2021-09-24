package main

import (
	"exaple.com/gRPC/grpctest"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000",grpc.WithInsecure(), grpc.WithBlock())
	if err!=nil{
		log.Fatalf("could not connect to server %s",err)
	}
	log.Printf("Connecting to server %s",conn.GetState())
	defer conn.Close()

	c:=grpctest.NewGrpctestClient(conn)

	message := grpctest.Message{
		Body: "Hello from the client!",
	}
	response, err := c.GrpcConnect(context.Background(),&message)
	if err!=nil{
		log.Fatalf("Error when speaking to Server %s",err)
	}

	log.Printf("Response from server %s",response.Body)
}

