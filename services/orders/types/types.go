package types

import (
	"context"
	"github.com/karmeev/common/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
