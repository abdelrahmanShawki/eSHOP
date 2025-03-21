package main

import (
	"github.com/abdelrahmanShawki/eSHOP/order/config"
	"github.com/abdelrahmanShawki/eSHOP/order/internal/adapter/db"
	"github.com/abdelrahmanShawki/eSHOP/order/internal/adapter/grpc"
	"github.com/abdelrahmanShawki/eSHOP/order/internal/application/core/service"
	"log"
)

func main() {
	log.Println("Starting application...")

	dataSourceURL := config.GetDataSourceURL()
	log.Println("Using Data Source URL:", dataSourceURL)

	dbAdapter, err := db.NewAdapter(dataSourceURL)
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	log.Println("Database connected successfully")

	orderService := service.NewOrderService(dbAdapter)
	grpcAdapter := grpc.NewAdapter(orderService, config.GetApplicationPort())

	log.Println("Starting gRPC server...")
	grpcAdapter.Run()

	log.Println("Application is running")
}
