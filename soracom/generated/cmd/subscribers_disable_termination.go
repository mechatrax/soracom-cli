// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SubscribersDisableTerminationCmdImsi holds value of 'imsi' option
var SubscribersDisableTerminationCmdImsi string

func init() {
	SubscribersDisableTerminationCmd.Flags().StringVar(&SubscribersDisableTerminationCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber."))
	SubscribersCmd.AddCommand(SubscribersDisableTerminationCmd)
}

// SubscribersDisableTerminationCmd defines 'disable-termination' subcommand
var SubscribersDisableTerminationCmd = &cobra.Command{
	Use:   "disable-termination",
	Short: TRAPI("/subscribers/{imsi}/disable_termination:post:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/disable_termination:post:description`),
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

		param, err := collectSubscribersDisableTerminationCmdParams(ac)
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

func collectSubscribersDisableTerminationCmdParams(ac *apiClient) (*apiParams, error) {
	if SubscribersDisableTerminationCmdImsi == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "imsi")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSubscribersDisableTerminationCmd("/subscribers/{imsi}/disable_termination"),
		query:  buildQueryForSubscribersDisableTerminationCmd(),
	}, nil
}

func buildPathForSubscribersDisableTerminationCmd(path string) string {

	escapedImsi := url.PathEscape(SubscribersDisableTerminationCmdImsi)

	path = strReplace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForSubscribersDisableTerminationCmd() url.Values {
	result := url.Values{}

	return result
}
