package main

import (
	"github.com/abdelrahmanShawki/eSHOP/order/config"
	"github.com/abdelrahmanShawki/eSHOP/order/internal/adapter/db"
	"github.com/abdelrahmanShawki/eSHOP/order/internal/adapter/grpc"
	"github.com/abdelrahmanShawki/eSHOP/order/internal/adapter/payment"
	"github.com/abdelrahmanShawki/eSHOP/order/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(paymentAdapter, dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
