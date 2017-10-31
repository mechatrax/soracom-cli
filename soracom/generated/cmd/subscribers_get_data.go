package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SubscribersGetDataCmdImsi holds value of 'imsi' option
var SubscribersGetDataCmdImsi string

// SubscribersGetDataCmdLastEvaluatedKey holds value of 'last_evaluated_key' option
var SubscribersGetDataCmdLastEvaluatedKey string

// SubscribersGetDataCmdSort holds value of 'sort' option
var SubscribersGetDataCmdSort string

// SubscribersGetDataCmdFrom holds value of 'from' option
var SubscribersGetDataCmdFrom int64

// SubscribersGetDataCmdLimit holds value of 'limit' option
var SubscribersGetDataCmdLimit int64

// SubscribersGetDataCmdTo holds value of 'to' option
var SubscribersGetDataCmdTo int64

func init() {
	SubscribersGetDataCmd.Flags().StringVar(&SubscribersGetDataCmdImsi, "imsi", "", TRAPI("IMSI of the target subscriber that generated data entries."))

	SubscribersGetDataCmd.Flags().StringVar(&SubscribersGetDataCmdLastEvaluatedKey, "last-evaluated-key", "", TRAPI("The value of `time` in the last log entry retrieved in the previous page. By specifying this parameter, you can continue to retrieve the list from the next page onward."))

	SubscribersGetDataCmd.Flags().StringVar(&SubscribersGetDataCmdSort, "sort", "", TRAPI("Sort order of the data entries. Either descending (latest data entry first) or ascending (oldest data entry first)."))

	SubscribersGetDataCmd.Flags().Int64Var(&SubscribersGetDataCmdFrom, "from", 0, TRAPI("Start time for the data entries search range (unixtime in milliseconds)."))

	SubscribersGetDataCmd.Flags().Int64Var(&SubscribersGetDataCmdLimit, "limit", 0, TRAPI("Maximum number of data entries to retrieve."))

	SubscribersGetDataCmd.Flags().Int64Var(&SubscribersGetDataCmdTo, "to", 0, TRAPI("End time for the data entries search range (unixtime in milliseconds)."))

	SubscribersCmd.AddCommand(SubscribersGetDataCmd)
}

// SubscribersGetDataCmd defines 'get-data' subcommand
var SubscribersGetDataCmd = &cobra.Command{
	Use:   "get-data",
	Short: TRAPI("/subscribers/{imsi}/data:get:summary"),
	Long:  TRAPI(`/subscribers/{imsi}/data:get:description`),
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

		param, err := collectSubscribersGetDataCmdParams()
		if err != nil {
			return err
		}

		result, err := ac.callAPI(param)
		if err != nil {
			cmd.SilenceUsage = true
			return err
		}

		if result == "" {
			return nil
		}

		return prettyPrintStringAsJSON(result)
	},
}

func collectSubscribersGetDataCmdParams() (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForSubscribersGetDataCmd("/subscribers/{imsi}/data"),
		query:  buildQueryForSubscribersGetDataCmd(),
	}, nil
}

func buildPathForSubscribersGetDataCmd(path string) string {

	path = strings.Replace(path, "{"+"imsi"+"}", SubscribersGetDataCmdImsi, -1)

	return path
}

func buildQueryForSubscribersGetDataCmd() string {
	result := []string{}

	if SubscribersGetDataCmdLastEvaluatedKey != "" {
		result = append(result, sprintf("%s=%s", "last_evaluated_key", SubscribersGetDataCmdLastEvaluatedKey))
	}

	if SubscribersGetDataCmdSort != "" {
		result = append(result, sprintf("%s=%s", "sort", SubscribersGetDataCmdSort))
	}

	if SubscribersGetDataCmdFrom != 0 {
		result = append(result, sprintf("%s=%d", "from", SubscribersGetDataCmdFrom))
	}

	if SubscribersGetDataCmdLimit != 0 {
		result = append(result, sprintf("%s=%d", "limit", SubscribersGetDataCmdLimit))
	}

	if SubscribersGetDataCmdTo != 0 {
		result = append(result, sprintf("%s=%d", "to", SubscribersGetDataCmdTo))
	}

	return strings.Join(result, "&")
}