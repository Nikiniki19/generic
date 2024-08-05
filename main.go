package main

import (
	"generic/proto"
	"generic/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8091"
)

func main() {

	grpcGenericConnection(port)

}
func grpcGenericConnection(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterGenericRequestServer(grpcServer, &service.FetchUser{})
	log.Printf("Server started at : %v", lis.Addr())

	err1 := grpcServer.Serve(lis)
	if err1 != nil {
		log.Fatalf("Failed to start: %v", err1)
	}
}
