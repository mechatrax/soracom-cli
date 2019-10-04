// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// OrdersGetCmdOrderId holds value of 'order_id' option
var OrdersGetCmdOrderId string

func init() {
	OrdersGetCmd.Flags().StringVar(&OrdersGetCmdOrderId, "order-id", "", TRAPI("order_id"))

	OrdersGetCmd.MarkFlagRequired("order-id")

	OrdersCmd.AddCommand(OrdersGetCmd)
}

// OrdersGetCmd defines 'get' subcommand
var OrdersGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/orders/{order_id}:get:summary"),
	Long:  TRAPI(`/orders/{order_id}:get:description`),
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

		param, err := collectOrdersGetCmdParams(ac)
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

func collectOrdersGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForOrdersGetCmd("/orders/{order_id}"),
		query:  buildQueryForOrdersGetCmd(),
	}, nil
}

func buildPathForOrdersGetCmd(path string) string {

	escapedOrderId := url.PathEscape(OrdersGetCmdOrderId)

	path = strReplace(path, "{"+"order_id"+"}", escapedOrderId, -1)

	return path
}

func buildQueryForOrdersGetCmd() url.Values {
	result := url.Values{}

	return result
}
