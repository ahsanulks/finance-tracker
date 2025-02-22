package entity

import "time"

type Transaction struct {
	date    time.Time
	amount  int64
	content string
}
