package handler

import (
	"errors"
	"financetracker/internal/entity"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const (
	yearMonthInputFormat = "200601" // YYYYMM format
)

var TransactionHistoryCmd = &cobra.Command{
	Use:     "fintrack <YYYYMM> <file-path>",
	Long:    "Fintrack is a command-line tool that reads financial records from a specified file\nand filters them based on the given period (YYYYMM).\nThis helps in analyzing financial data efficiently.",
	Example: "  fintrack 202403 data/transactions.csv",
	Args:    ValidateTransactionHistoryArgs,
	PreRunE: ValidateCsvFileExist,
}

func ValidateTransactionHistoryArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("missing arguments: required <YYYYMM> <file-path>")
	}

	if len(args) > 2 {
		return errors.New("too many arguments: expected only <YYYYMM> <file-path>")
	}

	return nil
}

func ValidateCsvFileExist(cmd *cobra.Command, args []string) error {
	if _, err := os.Stat(args[1]); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", args[1])
	}
	return nil
}

type TransactionHistoryCli struct{}

func (thc *TransactionHistoryCli) GenerateTransactionHistoryReport(cmd *cobra.Command, args []string) error {
	_, err := convertInputToTransactionPeriod(args[0])
	if err != nil {
		return err
	}
	return nil
}

func convertInputToTransactionPeriod(strDate string) (entity.TransactionPeriod, error) {
	_, err := time.ParseInLocation(yearMonthInputFormat, strDate, time.Local)
	if err != nil {
		return entity.TransactionPeriod{}, errors.New("invalid date format: must be YYYYMM")
	}
	return entity.TransactionPeriod{}, nil
}
