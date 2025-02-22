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
