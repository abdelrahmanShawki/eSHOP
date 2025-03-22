package api

import (
	"github.com/abdelrahmanShawki/eSHOP/order/internal/application/core/domain"
	"github.com/abdelrahmanShawki/eSHOP/order/internal/port"
)

type Application struct {
	db      port.DBPort
	payment port.PaymentPort
}

func NewApplication(payment port.PaymentPort, db port.DBPort) *Application {
	return &Application{
		payment: payment,
		db:      db,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}

	paymentErr := a.payment.Charge(&order)
	if paymentErr != nil {
		return domain.Order{}, paymentErr
	}

	return order, nil
}
