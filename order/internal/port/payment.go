package port

import "github.com/abdelrahmanShawki/eSHOP/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(order *domain.Order) error
}
