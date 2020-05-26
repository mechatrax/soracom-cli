// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// UsersAuthKeysListCmdOperatorId holds value of 'operator_id' option
var UsersAuthKeysListCmdOperatorId string

// UsersAuthKeysListCmdUserName holds value of 'user_name' option
var UsersAuthKeysListCmdUserName string

func init() {
	UsersAuthKeysListCmd.Flags().StringVar(&UsersAuthKeysListCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersAuthKeysListCmd.Flags().StringVar(&UsersAuthKeysListCmdUserName, "user-name", "", TRAPI("user_name"))
	UsersAuthKeysCmd.AddCommand(UsersAuthKeysListCmd)
}

// UsersAuthKeysListCmd defines 'list' subcommand
var UsersAuthKeysListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/auth_keys:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/auth_keys:get:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}
		err := authHelper(ac, cmd, args)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		param, err := collectUsersAuthKeysListCmdParams(ac)
		if err != nil {
			return err
		}

		body, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if body == "" {
			return nil
		}
		return prettyPrintStringAsJSON(body)

	},
}

func collectUsersAuthKeysListCmdParams(ac *apiClient) (*apiParams, error) {
	if UsersAuthKeysListCmdOperatorId == "" {
		UsersAuthKeysListCmdOperatorId = ac.OperatorID
	}

	if UsersAuthKeysListCmdUserName == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "user-name")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersAuthKeysListCmd("/operators/{operator_id}/users/{user_name}/auth_keys"),
		query:  buildQueryForUsersAuthKeysListCmd(),
	}, nil
}

func buildPathForUsersAuthKeysListCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersAuthKeysListCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersAuthKeysListCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersAuthKeysListCmd() url.Values {
	result := url.Values{}

	return result
}
