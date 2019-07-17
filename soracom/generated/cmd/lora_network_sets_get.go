// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraNetworkSetsGetCmdNsId holds value of 'ns_id' option
var LoraNetworkSetsGetCmdNsId string

func init() {
	LoraNetworkSetsGetCmd.Flags().StringVar(&LoraNetworkSetsGetCmdNsId, "ns-id", "", TRAPI("ID of the target LoRa network set."))

	LoraNetworkSetsGetCmd.MarkFlagRequired("ns-id")

	LoraNetworkSetsCmd.AddCommand(LoraNetworkSetsGetCmd)
}

// LoraNetworkSetsGetCmd defines 'get' subcommand
var LoraNetworkSetsGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/lora_network_sets/{ns_id}:get:summary"),
	Long:  TRAPI(`/lora_network_sets/{ns_id}:get:description`),
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

		param, err := collectLoraNetworkSetsGetCmdParams(ac)
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

func collectLoraNetworkSetsGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForLoraNetworkSetsGetCmd("/lora_network_sets/{ns_id}"),
		query:  buildQueryForLoraNetworkSetsGetCmd(),
	}, nil
}

func buildPathForLoraNetworkSetsGetCmd(path string) string {

	escapedNsId := url.PathEscape(LoraNetworkSetsGetCmdNsId)

	path = strings.Replace(path, "{"+"ns_id"+"}", escapedNsId, -1)

	return path
}

func buildQueryForLoraNetworkSetsGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
