// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(SystemNotificationsCmd)
}

// SystemNotificationsCmd defines 'system-notifications' subcommand
var SystemNotificationsCmd = &cobra.Command{
	Use:   "system-notifications",
	Short: TRCLI("cli.system-notifications.summary"),
	Long:  TRCLI(`cli.system-notifications.description`),
}
