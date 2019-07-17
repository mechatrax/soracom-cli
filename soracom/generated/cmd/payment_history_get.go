// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// PaymentHistoryGetCmdPaymentTransactionId holds value of 'payment_transaction_id' option
var PaymentHistoryGetCmdPaymentTransactionId string

func init() {
	PaymentHistoryGetCmd.Flags().StringVar(&PaymentHistoryGetCmdPaymentTransactionId, "payment-transaction-id", "", TRAPI("Payment transaction ID"))

	PaymentHistoryGetCmd.MarkFlagRequired("payment-transaction-id")

	PaymentHistoryCmd.AddCommand(PaymentHistoryGetCmd)
}

// PaymentHistoryGetCmd defines 'get' subcommand
var PaymentHistoryGetCmd = &cobra.Command{
	Use:   "get",
	Short: TRAPI("/payment_history/transactions/{payment_transaction_id}:get:summary"),
	Long:  TRAPI(`/payment_history/transactions/{payment_transaction_id}:get:description`),
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

		param, err := collectPaymentHistoryGetCmdParams(ac)
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

func collectPaymentHistoryGetCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForPaymentHistoryGetCmd("/payment_history/transactions/{payment_transaction_id}"),
		query:  buildQueryForPaymentHistoryGetCmd(),
	}, nil
}

func buildPathForPaymentHistoryGetCmd(path string) string {

	escapedPaymentTransactionId := url.PathEscape(PaymentHistoryGetCmdPaymentTransactionId)

	path = strings.Replace(path, "{"+"payment_transaction_id"+"}", escapedPaymentTransactionId, -1)

	return path
}

func buildQueryForPaymentHistoryGetCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}
