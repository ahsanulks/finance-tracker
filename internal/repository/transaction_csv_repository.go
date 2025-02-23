package repository

import (
	"context"
	"encoding/csv"
	"financetracker/internal/entity"
	"os"
	"slices"
	"strconv"
	"time"
)

const (
	csvTransactionDateFormat = "2006/01/02" //YYYY/MM/DD
)

type TransactionCsvRepository struct {
	filePath string
}

func NewTransactionCsvRepository(filePath string) *TransactionCsvRepository {
	return &TransactionCsvRepository{
		filePath: filePath,
	}
}

func (tcr *TransactionCsvRepository) FetchByPeriodDesc(
	ctx context.Context,
	transactionPeriod entity.TransactionPeriod,
) ([]*entity.Transaction, error) {
	file, err := os.Open(tcr.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// skip header
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	var transactions []*entity.Transaction
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		amount, _ := strconv.ParseInt(record[1], 10, 64)
		trxDate, _ := time.ParseInLocation(csvTransactionDateFormat, record[0], time.Local)
		if transactionPeriod.IsSamePeriod(trxDate) {
			transactions = append(transactions, entity.NewTransaction(trxDate, amount, record[2]))
		}
	}

	tcr.sortTransactionsDesc(transactions)

	return transactions, nil
}

func (tcr *TransactionCsvRepository) sortTransactionsDesc(transactions []*entity.Transaction) {
	slices.SortFunc(transactions, func(left, right *entity.Transaction) int {
		if left.Date().After(right.Date()) {
			return -1 // Descending order
		} else if left.Date().Before(right.Date()) {
			return 1
		}
		return 0
	})
}
