// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// StatsBeamGetCmdImsi holds value of 'imsi' option
var StatsBeamGetCmdImsi string

// StatsBeamGetCmdPeriod holds value of 'period' option
var StatsBeamGetCmdPeriod string

// StatsBeamGetCmdFrom holds value of 'from' option
var StatsBeamGetCmdFrom int64

// StatsBeamGetCmdTo holds value of 'to' option
var StatsBeamGetCmdTo int64

func init() {
	StatsBeamGetCmd.Flags().StringVar(&StatsBeamGetCmdImsi, "imsi", "", TRAPI("imsi"))

	StatsBeamGetCmd.MarkFlagRequired("imsi")

	StatsBeamGetCmd.Flags().StringVar(&StatsBeamGetCmdPeriod, "period", "", TRAPI("Units of aggregate data. For minutes, the interval is around 5 minutes."))

	StatsBeamGetCmd.MarkFlagRequired("period")

	StatsBeamGetCmd.Flags().Int64Var(&StatsBeamGetCmdFrom, "from", 0, TRAPI("Start time in unixtime for the aggregate data."))

	StatsBeamGetCmd.MarkFlagRequired("from")

	StatsBeamGetCmd.Flags().Int64Var(&StatsBeamGetCmdTo, "to", 0, TRAPI("End time in unixtime for the aggregate data."))

	StatsBeamGetCmd.MarkFlagRequired("to")

	StatsBeamCmd.AddCommand(StatsBeamGetCmd)
}

// StatsBeamGetCmd defines 'get' subcommand
var StatsBeamGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/stats/beam/subscribers/{imsi}:get:summary"),
	Long:  TRAPI(`/stats/beam/subscribers/{imsi}:get:description`),
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

		param, err := collectStatsBeamGetCmdParams(ac)
		if err != nil {
			return err
		}

		_, body, err := ac.callAPI(param)
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

func collectStatsBeamGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForStatsBeamGetCmd("/stats/beam/subscribers/{imsi}"),
		query:  buildQueryForStatsBeamGetCmd(),
	}, nil
}

func buildPathForStatsBeamGetCmd(path string) string {

	escapedImsi := url.PathEscape(StatsBeamGetCmdImsi)

	path = strings.Replace(path, "{"+"imsi"+"}", escapedImsi, -1)

	return path
}

func buildQueryForStatsBeamGetCmd() string {
	result := []string{}

	if StatsBeamGetCmdPeriod != "" {
		result = append(result, sprintf("%s=%s", url.QueryEscape("period"), url.QueryEscape(StatsBeamGetCmdPeriod)))
	}

	if StatsBeamGetCmdFrom != 0 {
		result = append(result, sprintf("%s=%s", url.QueryEscape("from"), url.QueryEscape(sprintf("%d", StatsBeamGetCmdFrom))))
	}

	if StatsBeamGetCmdTo != 0 {
		result = append(result, sprintf("%s=%s", url.QueryEscape("to"), url.QueryEscape(sprintf("%d", StatsBeamGetCmdTo))))
	}

	return strings.Join(result, "&")
}
