package repository_test

import (
	"context"
	"financetracker/internal/entity"
	"financetracker/internal/repository"
	"testing"
)

func TestTransactionHistoryStdoutRepository_Write(t *testing.T) {
	type args struct {
		ctx                context.Context
		transactionHistory *entity.TransactionHistory
	}
	tests := []struct {
		name    string
		thsr    *repository.TransactionHistoryStdoutRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thsr := &repository.TransactionHistoryStdoutRepository{}
			if err := thsr.Write(tt.args.ctx, tt.args.transactionHistory); (err != nil) != tt.wantErr {
				t.Errorf("TransactionHistoryStdoutRepository.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
