package usecase_test

import (
	"context"
	"financetracker/internal/usecase"
	"testing"
	"time"
)

func TestTransactionHistoryUsecase_GenerateHistoryByDate(t *testing.T) {
	type args struct {
		ctx  context.Context
		date time.Time
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
				ctx:  context.Background(),
				date: time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thu := &usecase.TransactionHistoryUsecase{}
			if err := thu.GenerateHistoryByDate(tt.args.ctx, tt.args.date); (err != nil) != tt.wantErr {
				t.Errorf("TransactionHistoryUsecase.GenerateHistoryByDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
