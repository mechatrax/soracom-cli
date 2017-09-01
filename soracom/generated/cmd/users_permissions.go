package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	UsersCmd.AddCommand(UsersPermissionsCmd)
}

// UsersPermissionsCmd defines 'permissions' subcommand
var UsersPermissionsCmd = &cobra.Command{
	Use:   "permissions",
	Short: TRCLI("cli.users.permissions.summary"),
	Long:  TRCLI(`cli.users.permissions.description`),
}
