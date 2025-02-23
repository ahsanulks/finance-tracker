package repository

import (
	"context"
	"financetracker/internal/entity"
)

type TransactionHistoryCsvRepository struct {
}

func (thc *TransactionHistoryCsvRepository) FetchByPeriodDesc(
	ctx context.Context,
	period entity.TransactionPeriod,
) ([]*entity.Transaction, error) {
	return nil, nil
}
