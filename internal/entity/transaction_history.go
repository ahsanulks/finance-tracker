package entity

import "errors"

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
