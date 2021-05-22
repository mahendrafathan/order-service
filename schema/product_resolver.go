package schema

import (
	"context"
	"strconv"

	"github.com/mahendrafathan/order-service/database"
)

func (r *Resolver) Products(ctx context.Context) ([]*database.Product, error) {
	return database.FindAllProductFunc()
}

func (r *Resolver) Product(ctx context.Context, args struct{ Id string }) (database.Product, error) {
	id, err := strconv.Atoi(args.Id)
	if err != nil {
		return database.Product{}, err
	}

	return database.GetProductByIdFunc(id)
}
