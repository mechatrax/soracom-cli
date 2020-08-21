// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// OperatorGetSupportTokenCmdOperatorId holds value of 'operator_id' option
var OperatorGetSupportTokenCmdOperatorId string

func init() {
	OperatorGetSupportTokenCmd.Flags().StringVar(&OperatorGetSupportTokenCmdOperatorId, "operator-id", "", TRAPI("operator_id"))
	OperatorCmd.AddCommand(OperatorGetSupportTokenCmd)
}

// OperatorGetSupportTokenCmd defines 'get-support-token' subcommand
var OperatorGetSupportTokenCmd = &cobra.Command{
	Use:   "get-support-token",
	Short: TRAPI("/operators/{operator_id}/support/token:post:summary"),
	Long:  TRAPI(`/operators/{operator_id}/support/token:post:description`),
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

		param, err := collectOperatorGetSupportTokenCmdParams(ac)
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

func collectOperatorGetSupportTokenCmdParams(ac *apiClient) (*apiParams, error) {
	if OperatorGetSupportTokenCmdOperatorId == "" {
		OperatorGetSupportTokenCmdOperatorId = ac.OperatorID
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForOperatorGetSupportTokenCmd("/operators/{operator_id}/support/token"),
		query:  buildQueryForOperatorGetSupportTokenCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForOperatorGetSupportTokenCmd(path string) string {

	escapedOperatorId := url.PathEscape(OperatorGetSupportTokenCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorGetSupportTokenCmd() url.Values {
	result := url.Values{}

	return result
}
