package repository

import (
	"context"
	"encoding/json"
	"financetracker/internal/entity"
	"os"
	"strconv"
	"time"
)

const (
	dailyFormatDateReport   = "2006/01/02" // format: YYYY/MM/DD
	monthlyFormatDateReport = "2006/01"    // format: YYYY/MM
)

type TransactionHistoryStdoutRepository struct{}

func (thsr *TransactionHistoryStdoutRepository) Write(
	ctx context.Context,
	transactionHistory *entity.TransactionHistory,
) error {
	transactionReport := NewTransactionReport(transactionHistory)
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(transactionReport)
}

type (
	DailyDateReport   time.Time
	MonthlyDateReport time.Time

	TransactionReport struct {
		Date         MonthlyDateReport `json:"period"`
		TotalIncome  int64             `json:"total_income"`
		TotalExpense int64             `json:"total_expense"`
		Transactions []*Transaction    `json:"transactions"`
	}

	Transaction struct {
		Date    DailyDateReport `json:"date"`
		Amount  string          `json:"amount"`
		Content string          `json:"content"`
	}
)

func NewTransactionReport(transactionHistory *entity.TransactionHistory) *TransactionReport {
	transactions := make([]*Transaction, 0, len(transactionHistory.Transactions()))
	for _, transaction := range transactionHistory.Transactions() {
		transactions = append(transactions, &Transaction{
			Date:    DailyDateReport(transaction.Date()),
			Amount:  strconv.FormatInt(transaction.Amount(), 10),
			Content: transaction.Content(),
		})
	}
	return &TransactionReport{
		Date:         MonthlyDateReport(time.Date(transactionHistory.YearPeriod(), time.Month(transactionHistory.MonthPeriod()), 1, 0, 0, 0, 0, time.Local)),
		TotalIncome:  transactionHistory.TotalIncome(),
		TotalExpense: transactionHistory.TotalExpenditure(),
		Transactions: transactions,
	}
}

func (ddr DailyDateReport) MarshalJSON() ([]byte, error) {
	t := time.Time(ddr)
	return json.Marshal(t.Format(dailyFormatDateReport))
}

func (mdr MonthlyDateReport) MarshalJSON() ([]byte, error) {
	t := time.Time(mdr)
	return json.Marshal(t.Format(monthlyFormatDateReport))
}
