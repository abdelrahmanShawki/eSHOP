package service

import (
	"github.com/abdelrahmanShawki/eSHOP/order/internal/application/core/domain"
	"github.com/abdelrahmanShawki/eSHOP/order/internal/port"
)

type OrderService struct {
	db port.DBPort
}

func NewOrderService(db port.DBPort) *OrderService {
	return &OrderService{
		db: db,
	}
}

func (s OrderService) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := s.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
