package grpctest
import (
	"golang.org/x/net/context"
	"log"
)

type Server struct{
}

func(s *Server) grpctest(ctx context.Context, message Message) (*Message,error){
	log.Printf("received message body fromm Client: %s", message.Body)
	return &Message{Body: "Hello from the Serverworld"},nil
}