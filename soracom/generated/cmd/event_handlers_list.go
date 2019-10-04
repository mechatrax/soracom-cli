// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// EventHandlersListCmdTarget holds value of 'target' option
var EventHandlersListCmdTarget string

func init() {
	EventHandlersListCmd.Flags().StringVar(&EventHandlersListCmdTarget, "target", "", TRAPI("target"))

	EventHandlersCmd.AddCommand(EventHandlersListCmd)
}

// EventHandlersListCmd defines 'list' subcommand
var EventHandlersListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/event_handlers:get:summary"),
	Long:  TRAPI(`/event_handlers:get:description`),
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

		param, err := collectEventHandlersListCmdParams(ac)
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

func collectEventHandlersListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForEventHandlersListCmd("/event_handlers"),
		query:  buildQueryForEventHandlersListCmd(),
	}, nil
}

func buildPathForEventHandlersListCmd(path string) string {

	return path
}

func buildQueryForEventHandlersListCmd() url.Values {
	result := url.Values{}

	if EventHandlersListCmdTarget != "" {
		result.Add("target", EventHandlersListCmdTarget)
	}

	return result
}
