// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(OrdersCmd)
}

// OrdersCmd defines 'orders' subcommand
var OrdersCmd = &cobra.Command{
	Use:   "orders",
	Short: TRCLI("cli.orders.summary"),
	Long:  TRCLI(`cli.orders.description`),
}
