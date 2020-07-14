// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// SimsGetCmdSimId holds value of 'sim_id' option
var SimsGetCmdSimId string

func init() {
	SimsGetCmd.Flags().StringVar(&SimsGetCmdSimId, "sim-id", "", TRAPI("Id of the target SIM"))
	SimsCmd.AddCommand(SimsGetCmd)
}

// SimsGetCmd defines 'get' subcommand
var SimsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/sims/{sim_id}:get:summary"),
	Long:  TRAPI(`/sims/{sim_id}:get:description`),
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

		param, err := collectSimsGetCmdParams(ac)
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

func collectSimsGetCmdParams(ac *apiClient) (*apiParams, error) {
	if SimsGetCmdSimId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "sim-id")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForSimsGetCmd("/sims/{sim_id}"),
		query:  buildQueryForSimsGetCmd(),
	}, nil
}

func buildPathForSimsGetCmd(path string) string {

	escapedSimId := url.PathEscape(SimsGetCmdSimId)

	path = strReplace(path, "{"+"sim_id"+"}", escapedSimId, -1)

	return path
}

func buildQueryForSimsGetCmd() url.Values {
	result := url.Values{}

	return result
}
