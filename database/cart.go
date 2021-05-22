package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	unpaid = "UNPAID"
	paid   = "PAID"
)

var (
	FindAllCartFunc               = findAllCart
	AddToCartFunc                 = addToCart
	FindCartByUserIDAndStatusFunc = findCartByUserIDAndStatus
)

type Cart struct {
	UserID         int32
	ProductID      int32
	Total          float64
	Quantity       int32
	FreeItems      int32
	TotalFreeItems int32
	Description    string
	Status         string
}

func addToCart(productID int32, userID int32, quantity int32) (cart Cart, err error) {
	product, err := GetProductByIdFunc(int(productID))
	if err != nil {
		return
	}

	// check stock
	if product.Quantity < quantity {
		err = fmt.Errorf("out of stock")
		return
	}

	// check promocode
	price, freeItems, totalFreeItems := checkPromoCode(product, int(quantity))
	cart = Cart{
		UserID:         userID,
		ProductID:      productID,
		Total:          price,
		Quantity:       quantity,
		FreeItems:      int32(freeItems),
		TotalFreeItems: int32(totalFreeItems),
	}
	err = insertCart(cart)

	return
}

func updateCartToPaid(cart Cart) error {
	query := `
	update cart set status = $1 where user_id = $2 and product_id  = $3
	`
	_, err := db.Exec(query, paid, cart.UserID, cart.ProductID)
	if err != nil {
		return err
	}

	return nil
}

func findCartByUserIDAndStatus(userID int32, status string) (carts []*Cart, err error) {
	query := `
	SELECT user_id, product_id, quantity, COALESCE(total, 0), coalesce(free_items, 0) ,
	coalesce(total_free_items, 0), coalesce(description, '') FROM cart where user_id = $1 and status = $2
	`

	rows, err := db.Query(query, userID, status)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var cart Cart
		err = rows.Scan(
			&cart.UserID,
			&cart.ProductID,
			&cart.Quantity,
			&cart.Total,
			&cart.FreeItems,
			&cart.TotalFreeItems,
			&cart.Description,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		carts = append(carts, &cart)
		log.Printf("%+v\n", cart)
	}

	return
}

func findAllCart() (carts []*Cart, err error) {
	query := `
	SELECT user_id, product_id, quantity, COALESCE(total, 0), coalesce(free_items, 0) ,
	coalesce(total_free_items, 0), coalesce(description, '') FROM cart
	`

	rows, err := db.Query(query)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var cart Cart
		err = rows.Scan(
			&cart.UserID,
			&cart.ProductID,
			&cart.Quantity,
			&cart.Total,
			&cart.FreeItems,
			&cart.TotalFreeItems,
			&cart.Description,
		)
		if err != nil {
			log.Println(err)
			continue
		}
		carts = append(carts, &cart)
	}

	return
}

func insertCart(cart Cart) error {
	if cart.Status == "" {
		cart.Status = unpaid
	}
	stmt, err := db.Prepare(`INSERT INTO cart (
		user_id, product_id, quantity, 
		total, free_items, total_free_items, status) values ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		cart.UserID,
		cart.ProductID,
		cart.Quantity,
		cart.Total,
		nullSqlInt(cart.FreeItems),
		nullSqlInt(cart.TotalFreeItems),
		cart.Status,
	)
	return err
}

func nullSqlInt(data int32) sql.NullInt32 {
	if data == 0 {
		return sql.NullInt32{}
	}

	return sql.NullInt32{
		Int32: data,
		Valid: true,
	}
}
