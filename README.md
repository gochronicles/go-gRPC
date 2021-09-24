# go-gRPC
Microservices communication over Remote Procedure Protocol

Download the above project --

Goto terminal and in the same directory that you created your server.go file and enable go modules

Your server is now ready to run

Now to create handlers for your .proto file, Fire -> protoc --go_out=. --go-grpc_out=. grpctest.proto

To register the proto services for the server to refer the grpctest.go file is present in the grpctest/

Now the client file is ready to run

Once server is running and Client is ready to run

You now have a gRPC enabled comunication between couple of services written in GoLang
