package entity

import "time"

type Order struct {
	ID           uint
	BuyerEmail   string
	BuyerAddress string
	ProductId    int
	OrderDate    time.Time
	Product      Product
}
