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
		args       args
		wantErr    bool
		wantResult *repository.TransactionReport
	}{
		{
			name: "should write transaction history in expected format",
			args: args{
				ctx:                context.Background(),
				transactionHistory: createValidTransactionHistory(),
			},
			wantErr: false,
			wantResult: &repository.TransactionReport{
				Date:         repository.MonthlyDateReport(time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC)),
				TotalIncome:  250,
				TotalExpense: -300,
				Transactions: []*repository.Transaction{
					{
						Date:    repository.DailyDateReport(time.Date(2024, 10, 30, 0, 0, 0, 0, time.UTC)),
						Amount:  "100",
						Content: "income",
					},
					{
						Date:    repository.DailyDateReport(time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC)),
						Amount:  "-250",
						Content: "buy stuff",
					},
					{
						Date:    repository.DailyDateReport(time.Date(2024, 10, 2, 0, 0, 0, 0, time.UTC)),
						Amount:  "-50",
						Content: "buy snack",
					},
					{
						Date:    repository.DailyDateReport(time.Date(2024, 10, 1, 0, 0, 0, 0, time.UTC)),
						Amount:  "150",
						Content: "income last month",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			oldStdout := os.Stdout // Save the original stdout
			readFile, writeFile, _ := os.Pipe()
			os.Stdout = writeFile
			defer func() {
				os.Stdout = oldStdout
			}()

			thsr := &repository.TransactionHistoryStdoutRepository{}
			if err := thsr.Write(tt.args.ctx, tt.args.transactionHistory); (err != nil) != tt.wantErr {
				t.Errorf("TransactionHistoryStdoutRepository.Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			_ = writeFile.Close()
			var buf bytes.Buffer
			_, _ = buf.ReadFrom(readFile)
			expectedOutput, _ := json.MarshalIndent(tt.wantResult, "", "  ")
			assert.Equal(string(expectedOutput)+"\n", buf.String())
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
