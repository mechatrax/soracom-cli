// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsDeactivateCmdSimId holds value of 'sim_id' option
var SimsDeactivateCmdSimId string

func init() {
	SimsDeactivateCmd.Flags().StringVar(&SimsDeactivateCmdSimId, "sim-id", "", TRAPI("SIM ID of the target SIM."))
	SimsCmd.AddCommand(SimsDeactivateCmd)
}

// SimsDeactivateCmd defines 'deactivate' subcommand
var SimsDeactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: TRAPI("/sims/{sim_id}/deactivate:post:summary"),
	Long:  TRAPI(`/sims/{sim_id}/deactivate:post:description`),
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

		param, err := collectSimsDeactivateCmdParams(ac)
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

func collectSimsDeactivateCmdParams(ac *apiClient) (*apiParams, error) {
	if SimsDeactivateCmdSimId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForSimsDeactivateCmd("/sims/{sim_id}/deactivate"),
		query:  buildQueryForSimsDeactivateCmd(),
	}, nil
}

func buildPathForSimsDeactivateCmd(path string) string {

	escapedSimId := url.PathEscape(SimsDeactivateCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsDeactivateCmd() url.Values {
	result := url.Values{}

	return result
}
