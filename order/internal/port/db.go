package port

import "github.com/abdelrahmanShawki/eSHOP/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
