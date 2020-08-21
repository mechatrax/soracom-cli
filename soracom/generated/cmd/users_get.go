// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// UsersGetCmdOperatorId holds value of 'operator_id' option
var UsersGetCmdOperatorId string

// UsersGetCmdUserName holds value of 'user_name' option
var UsersGetCmdUserName string

func init() {
	UsersGetCmd.Flags().StringVar(&UsersGetCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersGetCmd.Flags().StringVar(&UsersGetCmdUserName, "user-name", "", TRAPI("user_name"))
	UsersCmd.AddCommand(UsersGetCmd)
}

// UsersGetCmd defines 'get' subcommand
var UsersGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}:get:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}:get:description`),
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

		param, err := collectUsersGetCmdParams(ac)
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

		if rawOutput {
			_, err = os.Stdout.Write([]byte(body))
		} else {
			return prettyPrintStringAsJSON(body)
		}
		return err
	},
}

func collectUsersGetCmdParams(ac *apiClient) (*apiParams, error) {
	if UsersGetCmdOperatorId == "" {
		UsersGetCmdOperatorId = ac.OperatorID
	}

	if UsersGetCmdUserName == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "user-name")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForUsersGetCmd("/operators/{operator_id}/users/{user_name}"),
		query:  buildQueryForUsersGetCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForUsersGetCmd(path string) string {

	escapedOperatorId := url.PathEscape(UsersGetCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	escapedUserName := url.PathEscape(UsersGetCmdUserName)

	path = strReplace(path, "{"+"user_name"+"}", escapedUserName, -1)

	return path
}

func buildQueryForUsersGetCmd() url.Values {
	result := url.Values{}

	return result
}
