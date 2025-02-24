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
	}{
		{
			name: "when len args less than 2, should return error",
			args: args{
				cmd:  &cobra.Command{},
				args: []string{},
			},
			wantErr:    true,
			wantErrMsg: "missing arguments: required <YYYYMM> <file-path>",
		},
		{
			name: "when len args more than 2, should return error",
			args: args{
				cmd:  &cobra.Command{},
				args: []string{"202401", "file.csv", "test"},
			},
			wantErr:    true,
			wantErrMsg: "too many arguments: expected only <YYYYMM> <file-path>",
		},
		{
			name: "when date format not YYYYMM, shoudl return error",
			args: args{
				cmd:  &cobra.Command{},
				args: []string{"20241", "file.csv"},
			},
			wantErr:    true,
			wantErrMsg: "invalid date format: must be YYYYMM",
		},
		{
			name: "when input valid should return nil",
			args: args{
				cmd:  &cobra.Command{},
				args: []string{"202401", "file.csv"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := handler.TransactionHistoryCmd.Args(tt.args.cmd, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTransactionHistoryArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && tt.wantErrMsg != err.Error() {
				t.Errorf("ValidateTransactionHistoryArgs() errorMessage = %v, wantErrMessage %v", err, tt.wantErrMsg)
			}
		})
	}
}

func TestValidateCsvFileExist(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "when file not found, should return error",
			args: args{
				cmd:  &cobra.Command{},
				args: []string{"202401", "test.csv"},
			},
			wantErr:    true,
			wantErrMsg: "file not found: test.csv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handler.ValidateCsvFileExist(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCsvFileExist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
