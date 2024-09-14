package models

type Product struct {
	ID          int     `db:"id"`
	UserID      int     `db:"user_id"`
	Name        string  `db:"name" json:"name"`
	Description string  `db:"description" json:"description"`
	Price       float64 `db:"price" json:"price"`
	Quantity    int     `db:"quantity" json:"quantity"`
}
