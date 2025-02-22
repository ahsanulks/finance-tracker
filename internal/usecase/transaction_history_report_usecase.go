package usecase

import (
	"context"
	"financetracker/internal/entity"
)

type TransactionHistoryGetter interface {
	FetchByPeriod(ctx context.Context, period entity.TransactionPeriod) ([]*entity.Transaction, error)
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

func (thu *TransactionHistoryUsecase) GenerateHistoryByPeriod(
	ctx context.Context,
	period entity.TransactionPeriod,
) error {
	_, err := thu.transactionHistoryGetter.FetchByPeriod(ctx, period)
	if err != nil {
		return err
	}
	return nil
}
