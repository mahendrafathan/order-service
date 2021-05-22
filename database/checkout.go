package database

import "fmt"

var (
	CheckoutFunc = checkout
)

type Checkout struct {
	TotalOrder float64
	Message    string
	IsSuccess  bool
}

func checkout(userID int32, pay float64) Checkout {
	var data Checkout
	carts, err := FindCartByUserIDAndStatusFunc(userID, unpaid)
	if err != nil {
		data.Message = err.Error()
		return data
	}

	if len(carts) == 0 {
		data.Message = "Tidak Ada Barang di Keranjang"
		return data
	}

	var totalOrder float64
	for _, cart := range carts {
		totalOrder += cart.Total
	}

	if totalOrder != pay {
		data.Message = fmt.Sprintf("Nominal yang anda harus bayar : %+v", totalOrder)
		return data
	}

	// update stock
	for _, cart := range carts {
		err := updateStockProduct(cart.ProductID, cart.Quantity)
		if err != nil {
			continue
		}

		err = updateCartToPaid(*cart)
		if err != nil {
			continue
		}

	}

	return Checkout{
		TotalOrder: totalOrder,
		Message:    "Pembayaran berhasil",
		IsSuccess:  true,
	}
}
