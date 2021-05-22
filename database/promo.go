package database

import (
	"fmt"
)

type Promo struct {
	Code       string
	Discount   float64
	FreeItems  int64
	MinimumBuy int
}

func getPromoByID(code string) (promo Promo, err error) {
	query := `
		SELECT code, COALESCE(discount, 0), COALESCE(free_items, 0), minimum_buy
		FROM promo where code=$1
	`

	row := db.QueryRow(query, code)
	if row == nil {
		err = fmt.Errorf("row is nil")
		return
	}

	row.Scan(
		&promo.Code,
		&promo.Discount,
		&promo.FreeItems,
		&promo.MinimumBuy,
	)

	return
}
