package grpc

import (
	"fmt"
	"github.com/abdelrahmanShawki/eSHOP/order/config"
	"github.com/abdelrahmanShawki/eSHOP/order/internal/adapter/grpc/pb"
	"github.com/abdelrahmanShawki/eSHOP/order/internal/port"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Adapter struct {
	api  port.APIPort
	port int
	pb.UnimplementedOrderServer
}

func NewAdapter(api port.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	log.Println("Starting gRPC server setup...")

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d, error: %v", a.port, err)
	}
	log.Println("Successfully listening on port", a.port)

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServer(grpcServer, a)
	log.Println("Registered OrderServer service")

	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
		log.Println("Reflection enabled")
	}

	log.Println("Starting gRPC server on port", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
