package entity

import (
	"errors"
	"time"
)

type TransactionHistory struct {
	period           TransactionPeriod
	totalIncome      int64
	totalExpenditure int64
	transactions     []*Transaction
}

type TransactionPeriod struct {
	year  int
	month int
}

func NewTransactionPeriod(year int, month int) (TransactionPeriod, error) {
	if year < 1 {
		return TransactionPeriod{}, errors.New("year must be greater than 0")
	}
	if month < 1 || month > 12 {
		return TransactionPeriod{}, errors.New("month must be between 1 and 12")
	}
	return TransactionPeriod{year: year, month: month}, nil
}

func (tp TransactionPeriod) Year() int {
	return tp.year
}

func (tp TransactionPeriod) Month() int {
	return tp.month
}

func (tp TransactionPeriod) IsSamePeriod(transactionDate time.Time) bool {
	return tp.year == transactionDate.Year() && tp.month == int(transactionDate.Month())
}

func NewTransactionHistory(period TransactionPeriod, transactions []*Transaction) *TransactionHistory {
	transactionHistory := &TransactionHistory{
		period:       period,
		transactions: make([]*Transaction, len(transactions)),
	}

	for index, transaction := range transactions {
		transactionHistory.calculateTotal(transaction)
		transactionHistory.transactions[index] = transaction
	}

	return transactionHistory
}

func (th *TransactionHistory) calculateTotal(transaction *Transaction) {
	if transaction.IsExpense() {
		th.totalExpenditure += transaction.amount
	} else {
		th.totalIncome += transaction.amount
	}
}

func (th *TransactionHistory) YearPeriod() int {
	return th.period.year
}

func (th *TransactionHistory) MonthPeriod() int {
	return th.period.month
}

func (th *TransactionHistory) TotalIncome() int64 {
	return th.totalIncome
}

func (th *TransactionHistory) TotalExpenditure() int64 {
	return th.totalExpenditure
}

func (th *TransactionHistory) Transactions() []*Transaction {
	return th.transactions
}
