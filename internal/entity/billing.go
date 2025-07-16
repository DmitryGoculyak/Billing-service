package entity

import "time"

type Wallet struct {
	Id           string    `db:"id"`
	UserId       string    `db:"user_id"`
	CurrencyCode string    `db:"currency_code"`
	Balance      float64   `db:"balance"`
	CreatedAt    time.Time `db:"created_at"`
}
