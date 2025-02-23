package repository

import (
	"context"
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
	period entity.TransactionPeriod,
) ([]*entity.Transaction, error) {
	_, err := os.Open(thc.filePath)
	return nil, err
}
