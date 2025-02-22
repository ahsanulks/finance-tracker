package entity

import "time"

type TransactionHistory struct {
	period           time.Time
	totalIncome      int64
	totalExpenditure int64
	transactions     []*Transaction
}
