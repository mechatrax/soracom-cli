package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(PaymentHistoryCmd)
}

// PaymentHistoryCmd defines 'payment-history' subcommand
var PaymentHistoryCmd = &cobra.Command{
	Use:   "payment-history",
	Short: TRCLI("cli.payment-history.summary"),
	Long:  TRCLI(`cli.payment-history.description`),
}
