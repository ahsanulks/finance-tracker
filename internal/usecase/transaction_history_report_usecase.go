package usecase

import (
	"context"
	"financetracker/internal/entity"
)

type TransactionGetter interface {
	FetchByPeriodDesc(ctx context.Context, period entity.TransactionPeriod) ([]*entity.Transaction, error)
}

type TransactionHistoryWriter interface {
	Write(ctx context.Context, transactionHistory *entity.TransactionHistory) error
}

type TransactionHistoryUsecase struct {
	transactionGetter        TransactionGetter
	transactionHistoryWriter TransactionHistoryWriter
}

func NewTransactionHistoryUsecase(
	transactionGetter TransactionGetter,
	transactionHistoryWriter TransactionHistoryWriter,
) *TransactionHistoryUsecase {
	return &TransactionHistoryUsecase{
		transactionGetter:        transactionGetter,
		transactionHistoryWriter: transactionHistoryWriter,
	}
}

func (thu *TransactionHistoryUsecase) GenerateHistoryByPeriod(
	ctx context.Context,
	period entity.TransactionPeriod,
) error {
	transactions, err := thu.transactionGetter.FetchByPeriodDesc(ctx, period)
	if err != nil {
		return err
	}

	transactionHistory := entity.NewTransactionHistory(period, transactions)
	return thu.transactionHistoryWriter.Write(ctx, transactionHistory)
}
