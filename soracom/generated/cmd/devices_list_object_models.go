// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// DevicesListObjectModelsCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var DevicesListObjectModelsCmdLastEvaluatedKey string

// DevicesListObjectModelsCmdLimit holds value of 'limit' option
var DevicesListObjectModelsCmdLimit int64

// DevicesListObjectModelsCmdPaginate indicates to do pagination or not
var DevicesListObjectModelsCmdPaginate bool

func init() {
	DevicesListObjectModelsCmd.Flags().StringVar(&DevicesListObjectModelsCmdLastEvaluatedKey, "last-evaluated-key", "null", TRAPI("ID of the last device object model in the previous page"))

	DevicesListObjectModelsCmd.Flags().Int64Var(&DevicesListObjectModelsCmdLimit, "limit", -1, TRAPI("Max number of device object models in a response"))

	DevicesListObjectModelsCmd.Flags().BoolVar(&DevicesListObjectModelsCmdPaginate, "fetch-all", false, TRCLI("cli.common_params.paginate.short_help"))
	DevicesCmd.AddCommand(DevicesListObjectModelsCmd)
}

// DevicesListObjectModelsCmd defines 'list-object-models' subcommand
var DevicesListObjectModelsCmd = &cobra.Command{
	Use:   "list-object-models",
	Short: TRAPI("/device_object_models:get:summary"),
	Long:  TRAPI(`/device_object_models:get:description`),
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

		param, err := collectDevicesListObjectModelsCmdParams(ac)
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

func collectDevicesListObjectModelsCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForDevicesListObjectModelsCmd("/device_object_models"),
		query:  buildQueryForDevicesListObjectModelsCmd(),

		doPagination:                      DevicesListObjectModelsCmdPaginate,
		paginationKeyHeaderInResponse:     "x-soracom-next-key",
		paginationRequestParameterInQuery: "last_evaluated_key",
	}, nil
}

func buildPathForDevicesListObjectModelsCmd(path string) string {

	return path
}

func buildQueryForDevicesListObjectModelsCmd() url.Values {
	result := url.Values{}

	if DevicesListObjectModelsCmdLastEvaluatedKey != "null" {
		result.Add("last_evaluated_key", DevicesListObjectModelsCmdLastEvaluatedKey)
	}

	if DevicesListObjectModelsCmdLimit != -1 {
		result.Add("limit", sprintf("%d", DevicesListObjectModelsCmdLimit))
	}

	return result
}
