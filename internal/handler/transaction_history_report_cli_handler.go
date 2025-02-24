package handler

import (
	"github.com/spf13/cobra"
)

var TransactionHistoryCmd = &cobra.Command{
	Use:     "fintrack <YYYYMM> <file-path>",
	Long:    "Fintrack is a command-line tool that reads financial records from a specified file\nand filters them based on the given period (YYYYMM).\nThis helps in analyzing financial data efficiently.",
	Example: "  fintrack 202403 data/transactions.csv",
}
