package main

import (
	"log"
	"net"

	"github.com/nandanugg/BeliMangTestCasesPB2W4/caddy/entity/pb"
	"github.com/nandanugg/BeliMangTestCasesPB2W4/caddy/handler"
	"github.com/nandanugg/BeliMangTestCasesPB2W4/caddy/service"
	"google.golang.org/grpc"
)

func main() {
	locationService := service.NewMerchantService()

	handlers := handler.NewHandler(locationService)

	server := grpc.NewServer()
	pb.RegisterMerchantServiceServer(server, handlers)
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Server is running on port :50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
