package usecase_test

import (
	"context"
	"errors"
	"financetracker/internal/entity"
	"financetracker/internal/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransactionHistoryUsecase_GenerateHistoryByPeriod(t *testing.T) {
	type args struct {
		ctx    context.Context
		period entity.TransactionPeriod
	}

	transaction2025JanPeriod, _ := entity.NewTransactionPeriod(2025, 1)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "when failed get transaction history, should return error",
			args: args{
				ctx:    context.Background(),
				period: entity.TransactionPeriod{},
			},
			wantErr: true,
		},
		{
			name: "when success get transaction history, should write generate transaction history report",
			args: args{
				ctx:    context.Background(),
				period: transaction2025JanPeriod,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thu := usecase.NewTransactionHistoryUsecase(
				newFakeTransactionHistoryGetter(),
				newSpyTransactionHistoryWriter(t),
			)
			if err := thu.GenerateHistoryByPeriod(tt.args.ctx, tt.args.period); (err != nil) != tt.wantErr {
				t.Errorf("TransactionHistoryUsecase.GenerateHistoryByPeriod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

var _ usecase.TransactionHistoryGetter = new(FakeTransactionHistoryGetter)

type FakeTransactionHistoryGetter struct {
	transactions []*entity.Transaction
}

func newFakeTransactionHistoryGetter() *FakeTransactionHistoryGetter {
	return &FakeTransactionHistoryGetter{
		transactions: []*entity.Transaction{
			entity.NewTransaction(
				time.Date(2025, 1, 30, 0, 0, 0, 0, time.Local),
				100,
				"test transaction income",
			),
			entity.NewTransaction(
				time.Date(2025, 1, 29, 0, 0, 0, 0, time.Local),
				-100,
				"test transaction expense",
			),
			entity.NewTransaction(
				time.Date(2025, 1, 27, 0, 0, 0, 0, time.Local),
				500,
				"test transaction Income 2",
			),
			entity.NewTransaction(
				time.Date(2025, 1, 20, 0, 0, 0, 0, time.Local),
				-100,
				"test transaction expense 2",
			),
			entity.NewTransaction(
				time.Date(2025, 2, 1, 0, 0, 0, 0, time.Local),
				-100,
				"test transaction expense other period",
			),

			entity.NewTransaction(
				time.Date(2024, 2, 1, 0, 0, 0, 0, time.Local),
				100,
				"test transaction income other period",
			),
		},
	}
}

func (f *FakeTransactionHistoryGetter) FetchByPeriod(ctx context.Context, period entity.TransactionPeriod) ([]*entity.Transaction, error) {
	var result []*entity.Transaction
	for _, transaction := range f.transactions {
		if transaction.IsSamePeriod(period) {
			result = append(result, transaction)
		}
	}

	if len(result) == 0 {
		return nil, errors.New("error when fetch transaction history")
	}
	return result, nil
}

var _ usecase.TransactionHistoryWriter = new(SpyTransactionHistoryWriter)

type SpyTransactionHistoryWriter struct {
	assert *assert.Assertions
}

func newSpyTransactionHistoryWriter(t *testing.T) *SpyTransactionHistoryWriter {
	return &SpyTransactionHistoryWriter{
		assert: assert.New(t),
	}
}

func (s *SpyTransactionHistoryWriter) Write(ctx context.Context, transactionHistory *entity.TransactionHistory) error {
	return nil
}
