package repository_test

import (
	"bytes"
	"context"
	"encoding/json"
	"financetracker/internal/entity"
	"financetracker/internal/repository"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransactionHistoryStdoutRepository_Write(t *testing.T) {
	type args struct {
		ctx                context.Context
		transactionHistory *entity.TransactionHistory
	}
	tests := []struct {
		name       string
		thsr       *repository.TransactionHistoryStdoutRepository
		args       args
		wantErr    bool
		wantResult map[string]any
	}{
		{
			name: "should write transaction history in expected format",
			thsr: &repository.TransactionHistoryStdoutRepository{},
			args: args{
				ctx:                context.Background(),
				transactionHistory: createValidTransactionHistory(),
			},
			wantErr: false,
			wantResult: map[string]any{
				"period":            "2024/10",
				"total_income":      250,
				"total_expenditure": -300,
				"transactions": []map[string]any{
					{
						"date":    "2024/10/30",
						"amount":  "100",
						"content": "income",
					},
					{
						"date":    "2024/10/15",
						"amount":  "-250",
						"content": "buy stuff",
					},
					{
						"date":    "2024/10/2",
						"amount":  "-50",
						"content": "buy snack",
					},
					{
						"date":    "2024/10/1",
						"amount":  "150",
						"content": "income last month",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			oldStdout := os.Stdout // Save the original stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			thsr := &repository.TransactionHistoryStdoutRepository{}
			if err := thsr.Write(tt.args.ctx, tt.args.transactionHistory); (err != nil) != tt.wantErr {
				t.Errorf("TransactionHistoryStdoutRepository.Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			w.Close()
			var buf bytes.Buffer
			_, _ = buf.ReadFrom(r)
			os.Stdout = oldStdout
			expectedOutput, _ := json.Marshal(tt.wantResult)
			assert.Equal(string(expectedOutput), buf.String())
		})
	}
}

func createValidTransactionHistory() *entity.TransactionHistory {
	period, _ := entity.NewTransactionPeriod(2024, 10)
	return entity.NewTransactionHistory(
		period,
		[]*entity.Transaction{
			entity.NewTransaction(
				time.Date(2024, 10, 30, 0, 0, 0, 0, time.Local), 100, "income",
			),
			entity.NewTransaction(
				time.Date(2024, 10, 15, 0, 0, 0, 0, time.Local), -250, "buy stuff",
			),
			entity.NewTransaction(
				time.Date(2024, 10, 2, 0, 0, 0, 0, time.Local), -50, "buy snack",
			),
			entity.NewTransaction(
				time.Date(2024, 10, 1, 0, 0, 0, 0, time.Local), 150, "income last month",
			),
		},
	)
}
