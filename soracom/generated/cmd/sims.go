// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(SimsCmd)
}

// SimsCmd defines 'sims' subcommand
var SimsCmd = &cobra.Command{
	Use:   "sims",
	Short: TRCLI("cli.sims.summary"),
	Long:  TRCLI(`cli.sims.description`),
}
