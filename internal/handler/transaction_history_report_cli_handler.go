package handler

import (
	"errors"
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
}

func ValidateTransactionHistoryArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("missing arguments: required <YYYYMM> <file-path>")
	}

	if len(args) > 2 {
		return errors.New("too many arguments: expected only <YYYYMM> <file-path>")
	}

	_, err := time.ParseInLocation(yearMonthInputFormat, args[0], time.Local)
	if err != nil {
		return errors.New("invalid date format: must be YYYYMM")
	}
	return nil
}
