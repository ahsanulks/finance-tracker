package cli

import (
	"financetracker/internal/handler"

	"github.com/spf13/cobra"
)

func InitializeCliCommand() *cobra.Command {
	return handler.TransactionHistoryCmd
}
