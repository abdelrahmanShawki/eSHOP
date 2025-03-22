package grpc

import (
	"context"
	"github.com/abdelrahmanShawki/eSHOP/order/internal/adapter/grpc/pb"
	"github.com/abdelrahmanShawki/eSHOP/order/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, request *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	for _, orderItem := range request.Items {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{OrderId: result.ID}, nil
}
