package entity

import "time"

type Transaction struct {
	date    time.Time
	amount  int64
	content string
}

func NewTransaction(date time.Time, amount int64, content string) *Transaction {
	return &Transaction{
		date:    date,
		amount:  amount,
		content: content,
	}
}

func (t Transaction) IsExpense() bool {
	return t.amount < 0
}

func (t Transaction) IsSamePeriod(period TransactionPeriod) bool {
	return t.date.Year() == period.year && t.date.Month() == time.Month(period.month)
}
