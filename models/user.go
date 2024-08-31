package models

type User struct {
	ID           int    `db:"id"`
	Name         string `db:"name" json:"name"`
	Email        string `db:"email" json:"email"`
	Password     string `json:"password,omitempty"`
	PasswordHash string `db:"password_hash"`
	Role         string `db:"role" json:"role"`
}
