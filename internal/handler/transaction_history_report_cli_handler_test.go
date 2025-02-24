package handler_test

import (
	"financetracker/internal/handler"
	"testing"

	"github.com/spf13/cobra"
)

func TestValidateTransactionHistoryArgs(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := handler.ValidateTransactionHistoryArgs(tt.args.cmd, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTransactionHistoryArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && tt.wantErrMsg != err.Error() {
				t.Errorf("ValidateTransactionHistoryArgs() errorMessage = %v, wantErrMessage %v", err, tt.wantErrMsg)
			}
		})
	}
}
