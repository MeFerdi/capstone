package models

type Order struct {
	ID         int     `db:"id"`
	ProductID  int     `db:"product_id"`
	RetailerID int     `db:"retailer_id"`
	Quantity   int     `db:"quantity" json:"quantity"`
	TotalPrice float64 `db:"total_price"`
	Status     string  `db:"status" json:"status"`
}
