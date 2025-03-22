package port

import (
	"github.com/abdelrahmanShawki/eSHOP/order/internal/application/core/domain"
)

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
