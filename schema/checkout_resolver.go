package schema

import (
	"context"

	"github.com/mahendrafathan/order-service/database"
)

func (r *Resolver) Checkout(ctx context.Context, args struct {
	UserID int32
	Pay    float64
}) (database.Checkout, error) {
	checkout := database.CheckoutFunc(args.UserID, args.Pay)
	return checkout, nil
}
