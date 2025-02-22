package usecase_test

import (
	"context"
	"errors"
	"financetracker/internal/entity"
	"financetracker/internal/usecase"
	"testing"
)

func TestTransactionHistoryUsecase_GenerateHistoryByPeriod(t *testing.T) {
	type args struct {
		ctx    context.Context
		period entity.TransactionPeriod
	}
	tests := []struct {
		name    string
		thu     *usecase.TransactionHistoryUsecase
		args    args
		wantErr bool
	}{
		{
			name: "when failed get transaction history, should return error",
			thu:  &usecase.TransactionHistoryUsecase{},
			args: args{
				ctx:    context.Background(),
				period: entity.TransactionPeriod{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thu := usecase.NewTransactionHistoryUsecase(
				new(FakeTransactionHistoryGetter),
			)
			if err := thu.GenerateHistoryByPeriod(tt.args.ctx, tt.args.period); (err != nil) != tt.wantErr {
				t.Errorf("TransactionHistoryUsecase.GenerateHistoryByPeriod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

var _ usecase.TransactionHistoryGetter = new(FakeTransactionHistoryGetter)

type FakeTransactionHistoryGetter struct{}

func (f *FakeTransactionHistoryGetter) FetchByPeriod(ctx context.Context, period entity.TransactionPeriod) ([]*entity.Transaction, error) {
	return nil, errors.New("error when fetch transaction history")
}
