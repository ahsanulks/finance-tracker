package repository

import (
	"context"
	"financetracker/internal/entity"
)

type TransactionHistoryStdoutRepository struct{}

func (thsr *TransactionHistoryStdoutRepository) Write(
	ctx context.Context,
	transactionHistory *entity.TransactionHistory,
) error {
	return nil
}
