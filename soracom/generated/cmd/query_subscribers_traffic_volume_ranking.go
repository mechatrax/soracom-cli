// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// QuerySubscribersTrafficVolumeRankingCmdOrder holds value of 'order' option
var QuerySubscribersTrafficVolumeRankingCmdOrder string

// QuerySubscribersTrafficVolumeRankingCmdFrom holds value of 'from' option
var QuerySubscribersTrafficVolumeRankingCmdFrom int64

// QuerySubscribersTrafficVolumeRankingCmdLimit holds value of 'limit' option
var QuerySubscribersTrafficVolumeRankingCmdLimit int64

// QuerySubscribersTrafficVolumeRankingCmdTo holds value of 'to' option
var QuerySubscribersTrafficVolumeRankingCmdTo int64

func init() {
	QuerySubscribersTrafficVolumeRankingCmd.Flags().StringVar(&QuerySubscribersTrafficVolumeRankingCmdOrder, "order", "desc", TRAPI("The order of ranking"))

	QuerySubscribersTrafficVolumeRankingCmd.Flags().Int64Var(&QuerySubscribersTrafficVolumeRankingCmdFrom, "from", 0, TRAPI("The beginning point of searching range (unixtime: in milliseconds)"))

	QuerySubscribersTrafficVolumeRankingCmd.Flags().Int64Var(&QuerySubscribersTrafficVolumeRankingCmdLimit, "limit", 10, TRAPI("The maximum number of item to retrieve"))

	QuerySubscribersTrafficVolumeRankingCmd.Flags().Int64Var(&QuerySubscribersTrafficVolumeRankingCmdTo, "to", 0, TRAPI("The end point of searching range (unixtime: in milliseconds)"))
	QuerySubscribersCmd.AddCommand(QuerySubscribersTrafficVolumeRankingCmd)
}

// QuerySubscribersTrafficVolumeRankingCmd defines 'traffic-volume-ranking' subcommand
var QuerySubscribersTrafficVolumeRankingCmd = &cobra.Command{
	Use:   "traffic-volume-ranking",
	Short: TRAPI("/query/subscribers/traffic_volume/ranking:get:summary"),
	Long:  TRAPI(`/query/subscribers/traffic_volume/ranking:get:description`),
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

		param, err := collectQuerySubscribersTrafficVolumeRankingCmdParams(ac)
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

func collectQuerySubscribersTrafficVolumeRankingCmdParams(ac *apiClient) (*apiParams, error) {

	if QuerySubscribersTrafficVolumeRankingCmdFrom == 0 {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "from")
	}

	if QuerySubscribersTrafficVolumeRankingCmdTo == 0 {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "to")
	}

	return &apiParams{
		method: "GET",
		path:   buildPathForQuerySubscribersTrafficVolumeRankingCmd("/query/subscribers/traffic_volume/ranking"),
		query:  buildQueryForQuerySubscribersTrafficVolumeRankingCmd(),
	}, nil
}

func buildPathForQuerySubscribersTrafficVolumeRankingCmd(path string) string {

	return path
}

func buildQueryForQuerySubscribersTrafficVolumeRankingCmd() url.Values {
	result := url.Values{}

	if QuerySubscribersTrafficVolumeRankingCmdOrder != "desc" {
		result.Add("order", QuerySubscribersTrafficVolumeRankingCmdOrder)
	}

	if QuerySubscribersTrafficVolumeRankingCmdFrom != 0 {
		result.Add("from", sprintf("%d", QuerySubscribersTrafficVolumeRankingCmdFrom))
	}

	if QuerySubscribersTrafficVolumeRankingCmdLimit != 10 {
		result.Add("limit", sprintf("%d", QuerySubscribersTrafficVolumeRankingCmdLimit))
	}

	if QuerySubscribersTrafficVolumeRankingCmdTo != 0 {
		result.Add("to", sprintf("%d", QuerySubscribersTrafficVolumeRankingCmdTo))
	}

	return result
}
