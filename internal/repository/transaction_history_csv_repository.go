package repository

import (
	"context"
	"financetracker/internal/entity"
	"os"
)

type TransactionHistoryCsvRepository struct {
	filePath string
}

func NewTransactionHistoryCsvRepository(filePath string) *TransactionHistoryCsvRepository {
	return &TransactionHistoryCsvRepository{
		filePath: filePath,
	}
}

func (thc *TransactionHistoryCsvRepository) FetchByPeriodDesc(
	ctx context.Context,
	period entity.TransactionPeriod,
) ([]*entity.Transaction, error) {
	_, err := os.Open(thc.filePath)
	return nil, err
}
