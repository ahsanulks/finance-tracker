package repository_test

import (
	"context"
	"financetracker/internal/entity"
	"financetracker/internal/repository"
	"reflect"
	"testing"
	"time"
)

func TestTransactionCsvRepository_FetchByPeriodDesc(t *testing.T) {
	transactionPeriod, _ := entity.NewTransactionPeriod(2022, 1)
	emptyTransactionPeriod, _ := entity.NewTransactionPeriod(2025, 1)
	type args struct {
		ctx      context.Context
		period   entity.TransactionPeriod
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []*entity.Transaction
		wantErr bool
	}{
		{
			name: "when failed to read file, should return error",
			args: args{
				ctx:      context.Background(),
				period:   transactionPeriod,
				filePath: "test123.csv",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when failed to read header csv, should return error",
			args: args{
				ctx:      context.Background(),
				period:   transactionPeriod,
				filePath: "test_data/empty_data.csv",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "when read valid format, should return filtered transaction with desc order",
			args: args{
				ctx:      context.Background(),
				period:   transactionPeriod,
				filePath: "test_data/valid_format.csv",
			},
			want: []*entity.Transaction{
				entity.NewTransaction(
					time.Date(2022, 1, 25, 0, 0, 0, 0, time.Local), -100000, "rent",
				),
				entity.NewTransaction(
					time.Date(2022, 1, 20, 0, 0, 0, 0, time.Local), 1000, "cash back",
				),
				entity.NewTransaction(
					time.Date(2022, 1, 6, 0, 0, 0, 0, time.Local), -10000, "debit",
				),
				entity.NewTransaction(
					time.Date(2022, 1, 5, 0, 0, 0, 0, time.Local), -1000, "eating out",
				),
				entity.NewTransaction(
					time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local), 1000, "salary",
				),
			},
			wantErr: false,
		},
		{
			name: "when there's no transaction that have same period, should return error",
			args: args{
				ctx:      context.Background(),
				period:   emptyTransactionPeriod,
				filePath: "test_data/valid_format.csv",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thc := repository.NewTransactionCsvRepository(tt.args.filePath)
			got, err := thc.FetchByPeriodDesc(tt.args.ctx, tt.args.period)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionHistoryCsv.FetchByPeriodDesc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionHistoryCsv.FetchByPeriodDesc() = %v, want %v", got, tt.want)
			}
		})
	}
}
