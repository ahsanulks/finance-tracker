package usecase

import (
	"context"
	"time"
)

type TransactionHistoryUsecase struct {
}

func (thu *TransactionHistoryUsecase) GenerateHistoryByDate(
	ctx context.Context,
	date time.Time,
) error {
	return nil
}
