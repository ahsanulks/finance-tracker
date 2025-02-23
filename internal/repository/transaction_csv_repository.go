package repository

import (
	"context"
	"encoding/csv"
	"financetracker/internal/entity"
	"os"
)

type TransactionCsvRepository struct {
	filePath string
}

func NewTransactionCsvRepository(filePath string) *TransactionCsvRepository {
	return &TransactionCsvRepository{
		filePath: filePath,
	}
}

func (thc *TransactionCsvRepository) FetchByPeriodDesc(
	ctx context.Context,
	transactionPeriod entity.TransactionPeriod,
) ([]*entity.Transaction, error) {
	file, err := os.Open(thc.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// skip header
	_, err = reader.Read()
	return nil, err
}
