package repository_test

import (
	"context"
	"financetracker/internal/entity"
	"financetracker/internal/repository"
	"reflect"
	"testing"
)

func TestTransactionHistoryCsvRepository_FetchByPeriodDesc(t *testing.T) {
	transactionPeriod, _ := entity.NewTransactionPeriod(2025, 2)
	type args struct {
		ctx    context.Context
		period entity.TransactionPeriod
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
				ctx:    context.Background(),
				period: transactionPeriod,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thc := &repository.TransactionHistoryCsvRepository{}
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
