package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(LoraNetworkSetsCmd)
}

// LoraNetworkSetsCmd defines 'lora-network-sets' subcommand
var LoraNetworkSetsCmd = &cobra.Command{
	Use:   "lora-network-sets",
	Short: TRCLI("cli.lora-network-sets.summary"),
	Long:  TRCLI(`cli.lora-network-sets.description`),
}
