package main

import (
	handler "github.com/karmeev/common/services/orders/handler/orders"
	"github.com/karmeev/common/services/orders/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	list, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// register our gRPC services
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)

	log.Printf("gRPC server listening at %v", s.addr)

	return grpcServer.Serve(list)
}
