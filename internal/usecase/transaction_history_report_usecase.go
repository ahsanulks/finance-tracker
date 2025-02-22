package usecase

import (
	"context"
	"financetracker/internal/entity"
	"time"
)

type TransactionHistoryGetter interface {
	FetchByPeriod(ctx context.Context, period time.Time) ([]*entity.Transaction, error)
}

type TransactionHistoryUsecase struct {
	transactionHistoryGetter TransactionHistoryGetter
}

func NewTransactionHistoryUsecase(
	transactionHistoryGetter TransactionHistoryGetter,
) *TransactionHistoryUsecase {
	return &TransactionHistoryUsecase{
		transactionHistoryGetter: transactionHistoryGetter,
	}
}

func (thu *TransactionHistoryUsecase) GenerateHistoryByDate(
	ctx context.Context,
	date time.Time,
) error {
	_, err := thu.transactionHistoryGetter.FetchByPeriod(ctx, date)
	if err != nil {
		return err
	}
	return nil
}
