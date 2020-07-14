// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"io/ioutil"

	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// GadgetsPutTagsCmdProductId holds value of 'product_id' option
var GadgetsPutTagsCmdProductId string

// GadgetsPutTagsCmdSerialNumber holds value of 'serial_number' option
var GadgetsPutTagsCmdSerialNumber string

// GadgetsPutTagsCmdBody holds contents of request body to be sent
var GadgetsPutTagsCmdBody string

func init() {
	GadgetsPutTagsCmd.Flags().StringVar(&GadgetsPutTagsCmdProductId, "product-id", "", TRAPI("Product ID of the target gadget."))

	GadgetsPutTagsCmd.Flags().StringVar(&GadgetsPutTagsCmdSerialNumber, "serial-number", "", TRAPI("Serial Number of the target gadget."))

	GadgetsPutTagsCmd.Flags().StringVar(&GadgetsPutTagsCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	GadgetsCmd.AddCommand(GadgetsPutTagsCmd)
}

// GadgetsPutTagsCmd defines 'put-tags' subcommand
var GadgetsPutTagsCmd = &cobra.Command{
	Use:   "put-tags",
	Short: TRAPI("/gadgets/{product_id}/{serial_number}/tags:put:summary"),
	Long:  TRAPI(`/gadgets/{product_id}/{serial_number}/tags:put:description`),
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

		param, err := collectGadgetsPutTagsCmdParams(ac)
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

func collectGadgetsPutTagsCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForGadgetsPutTagsCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if GadgetsPutTagsCmdProductId == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "product-id")
		}

	}

	if GadgetsPutTagsCmdSerialNumber == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "serial-number")
		}

	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForGadgetsPutTagsCmd("/gadgets/{product_id}/{serial_number}/tags"),
		query:       buildQueryForGadgetsPutTagsCmd(),
		contentType: contentType,
		body:        body,
	}, nil
}

func buildPathForGadgetsPutTagsCmd(path string) string {

	escapedProductId := url.PathEscape(GadgetsPutTagsCmdProductId)

	path = strReplace(path, "{"+"product_id"+"}", escapedProductId, -1)

	escapedSerialNumber := url.PathEscape(GadgetsPutTagsCmdSerialNumber)

	path = strReplace(path, "{"+"serial_number"+"}", escapedSerialNumber, -1)

	return path
}

func buildQueryForGadgetsPutTagsCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForGadgetsPutTagsCmd() (string, error) {
	var b []byte
	var err error

	if GadgetsPutTagsCmdBody != "" {
		if strings.HasPrefix(GadgetsPutTagsCmdBody, "@") {
			fname := strings.TrimPrefix(GadgetsPutTagsCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if GadgetsPutTagsCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(GadgetsPutTagsCmdBody)
		}

		if err != nil {
			return "", err
		}
	}

	if b == nil {
		b = []byte{}
	}

	return string(b), nil
}
