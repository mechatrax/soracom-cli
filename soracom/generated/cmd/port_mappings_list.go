// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// PortMappingsListCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var PortMappingsListCmdLastEvaluatedKey string

// PortMappingsListCmdLimit holds value of 'limit' option
var PortMappingsListCmdLimit int64

// PortMappingsListCmdPaginate indicates to do pagination or not
var PortMappingsListCmdPaginate bool

func init() {
	PortMappingsListCmd.Flags().StringVar(&PortMappingsListCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The last Port Mapping ID retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next group onward."))

	PortMappingsListCmd.Flags().Int64Var(&PortMappingsListCmdLimit, "limit", 0, TRAPI("Maximum number of results per response page."))

	PortMappingsListCmd.Flags().BoolVar(&PortMappingsListCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	PortMappingsCmd.AddCommand(PortMappingsListCmd)
}

// PortMappingsListCmd defines 'list' subcommand
var PortMappingsListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/port_mappings:get:summary"),
	Long:  TRAPI(`/port_mappings:get:description`),
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

		param, err := collectPortMappingsListCmdParams(ac)
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

func collectPortMappingsListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForPortMappingsListCmd("/port_mappings"),
		query:  buildQueryForPortMappingsListCmd(),

		doPagination:                      PortMappingsListCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",
	}, nil
}

func buildPathForPortMappingsListCmd(path string) string {

	return path
}

func buildQueryForPortMappingsListCmd() url.Values {
	result := url.Values{}

	if PortMappingsListCmdLastEvaluatedKey != "" {
		result.Add("last_evaluated_key", PortMappingsListCmdLastEvaluatedKey)
	}

	if PortMappingsListCmdLimit != 0 {
		result.Add("limit", sprintf("%d", PortMappingsListCmdLimit))
	}

	return result
}
