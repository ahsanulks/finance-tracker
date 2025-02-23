package usecase

import (
	"context"
	"financetracker/internal/entity"
)

type TransactionHistoryGetter interface {
	FetchByPeriod(ctx context.Context, period entity.TransactionPeriod) ([]*entity.Transaction, error)
}

type TransactionHistoryWriter interface {
	Write(ctx context.Context, transactionHistory *entity.TransactionHistory) error
}

type TransactionHistoryUsecase struct {
	transactionHistoryGetter TransactionHistoryGetter
	transactionHistoryWriter TransactionHistoryWriter
}

func NewTransactionHistoryUsecase(
	transactionHistoryGetter TransactionHistoryGetter,
	transactionHistoryWriter TransactionHistoryWriter,
) *TransactionHistoryUsecase {
	return &TransactionHistoryUsecase{
		transactionHistoryGetter: transactionHistoryGetter,
		transactionHistoryWriter: transactionHistoryWriter,
	}
}

func (thu *TransactionHistoryUsecase) GenerateHistoryByPeriod(
	ctx context.Context,
	period entity.TransactionPeriod,
) error {
	transactions, err := thu.transactionHistoryGetter.FetchByPeriod(ctx, period)
	if err != nil {
		return err
	}

	transactionHistory := entity.NewTransactionHistory(period, transactions)
	return thu.transactionHistoryWriter.Write(ctx, transactionHistory)
}
