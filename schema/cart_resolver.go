package schema

import (
	"context"

	"github.com/mahendrafathan/order-service/database"
)

func (r *Resolver) Carts(ctx context.Context) ([]*database.Cart, error) {
	return database.FindAllCartFunc()
}

func (r *Resolver) AddToCart(ctx context.Context, args struct {
	ProductID int32
	UserID    int32
	Quantity  int32
}) (database.Cart, error) {
	return database.AddToCartFunc(args.ProductID, args.UserID, args.Quantity)
}
