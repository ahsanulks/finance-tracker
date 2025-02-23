package entity

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
