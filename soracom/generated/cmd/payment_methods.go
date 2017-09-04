package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(PaymentMethodsCmd)
}

// PaymentMethodsCmd defines 'payment-methods' subcommand
var PaymentMethodsCmd = &cobra.Command{
	Use:   "payment-methods",
	Short: TRCLI("cli.payment-methods.summary"),
	Long:  TRCLI(`cli.payment-methods.description`),
}
