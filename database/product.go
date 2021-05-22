package database

import (
	"fmt"
	"log"
)

var (
	FindAllProductFunc = findAllProduct
	GetProductByIdFunc = getProductById
)

type Product struct {
	Sku       string  `json:"sku"`
	Name      string  `json:"product_name"`
	Price     float64 `json:"price"`
	Quantity  int32   `json:"quantity"`
	PromoCode string  `json:"promo_code,omitempty"`
}

func findAllProduct() (products []*Product, err error) {
	query := `
		SELECT sku, product_name, price, quantity, COALESCE(promo_code, '')
		FROM product
	`

	rows, err := db.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var product Product
		err = rows.Scan(
			&product.Sku,
			&product.Name,
			&product.Price,
			&product.Quantity,
			&product.PromoCode)
		if err != nil {
			log.Println(err)
			continue
		}
		products = append(products, &product)
	}

	return
}

func getProductById(productID int) (product Product, err error) {
	query := `
		SELECT sku, product_name, price, quantity, COALESCE(promo_code, '')
		FROM product where id=$1
	`

	row := db.QueryRow(query, productID)
	if row == nil {
		err = fmt.Errorf("row is nil")
		return
	}

	row.Scan(
		&product.Sku,
		&product.Name,
		&product.Price,
		&product.Quantity,
		&product.PromoCode,
	)

	return
}

func updateStockProduct(productID int32, sold int32) error {
	product, err := getProductById(int(productID))
	if err != nil {
		return err
	}

	quantity := product.Quantity - sold
	query := `
		update product set quantity = $1 where id = $2
	`
	_, err = db.Exec(query, quantity, productID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func checkPromoCode(product Product, quantity int) (price float64, freeItems int64, totalFreeItems int) {
	promo, err := getPromoByID(product.PromoCode)
	if err != nil {

	}
	switch product.PromoCode {
	case "PROMO1":
		// TODO: check stock for free items
		freeItems = promo.FreeItems
		totalFreeItems = int(quantity / promo.MinimumBuy)
		price = product.Price * float64(quantity)
		return
	case "PROMO2":
		count := int(quantity / promo.MinimumBuy)
		quantity = quantity - count
		price = (product.Price * float64(quantity))
		return
	case "PROMO3":
		var disc float64
		total := (product.Price * float64(quantity))
		if quantity > promo.MinimumBuy {
			disc = total * promo.Discount / 100
		}
		price = total - disc
		return
	default:
		price = (product.Price * float64(quantity))
	}

	return
}
