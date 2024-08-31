package models

import "time"

type CartItem struct {
	ID         int       `db:"id"`
	RetailerID int       `db:"retailer_id"`
	ProductID  int       `db:"product_id"`
	Quantity   int       `db:"quantity"`
	AddedAt    time.Time `db:"added_at"`
}
