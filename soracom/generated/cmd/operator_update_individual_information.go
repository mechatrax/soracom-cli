// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

	"fmt"

	"io/ioutil"

	"net/url"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

// OperatorUpdateIndividualInformationCmdAddressLine1 holds value of 'addressLine1' option
var OperatorUpdateIndividualInformationCmdAddressLine1 string

// OperatorUpdateIndividualInformationCmdAddressLine2 holds value of 'addressLine2' option
var OperatorUpdateIndividualInformationCmdAddressLine2 string

// OperatorUpdateIndividualInformationCmdBuilding holds value of 'building' option
var OperatorUpdateIndividualInformationCmdBuilding string

// OperatorUpdateIndividualInformationCmdCity holds value of 'city' option
var OperatorUpdateIndividualInformationCmdCity string

// OperatorUpdateIndividualInformationCmdCountryCode holds value of 'countryCode' option
var OperatorUpdateIndividualInformationCmdCountryCode string

// OperatorUpdateIndividualInformationCmdFullName holds value of 'fullName' option
var OperatorUpdateIndividualInformationCmdFullName string

// OperatorUpdateIndividualInformationCmdOperatorId holds value of 'operator_id' option
var OperatorUpdateIndividualInformationCmdOperatorId string

// OperatorUpdateIndividualInformationCmdPhoneNumber holds value of 'phoneNumber' option
var OperatorUpdateIndividualInformationCmdPhoneNumber string

// OperatorUpdateIndividualInformationCmdState holds value of 'state' option
var OperatorUpdateIndividualInformationCmdState string

// OperatorUpdateIndividualInformationCmdZipCode holds value of 'zipCode' option
var OperatorUpdateIndividualInformationCmdZipCode string

// OperatorUpdateIndividualInformationCmdBody holds contents of request body to be sent
var OperatorUpdateIndividualInformationCmdBody string

func init() {
	OperatorUpdateIndividualInformationCmd.Flags().StringVar(&OperatorUpdateIndividualInformationCmdAddressLine1, "address-line1", "", TRAPI(""))

	OperatorUpdateIndividualInformationCmd.Flags().StringVar(&OperatorUpdateIndividualInformationCmdAddressLine2, "address-line2", "", TRAPI(""))

	OperatorUpdateIndividualInformationCmd.Flags().StringVar(&OperatorUpdateIndividualInformationCmdBuilding, "building", "", TRAPI(""))

	OperatorUpdateIndividualInformationCmd.Flags().StringVar(&OperatorUpdateIndividualInformationCmdCity, "city", "", TRAPI(""))

	OperatorUpdateIndividualInformationCmd.Flags().StringVar(&OperatorUpdateIndividualInformationCmdCountryCode, "country-code", "", TRAPI(""))

	OperatorUpdateIndividualInformationCmd.Flags().StringVar(&OperatorUpdateIndividualInformationCmdFullName, "full-name", "", TRAPI(""))

	OperatorUpdateIndividualInformationCmd.Flags().StringVar(&OperatorUpdateIndividualInformationCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	OperatorUpdateIndividualInformationCmd.Flags().StringVar(&OperatorUpdateIndividualInformationCmdPhoneNumber, "phone-number", "", TRAPI(""))

	OperatorUpdateIndividualInformationCmd.Flags().StringVar(&OperatorUpdateIndividualInformationCmdState, "state", "", TRAPI(""))

	OperatorUpdateIndividualInformationCmd.Flags().StringVar(&OperatorUpdateIndividualInformationCmdZipCode, "zip-code", "", TRAPI(""))

	OperatorUpdateIndividualInformationCmd.Flags().StringVar(&OperatorUpdateIndividualInformationCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	OperatorCmd.AddCommand(OperatorUpdateIndividualInformationCmd)
}

// OperatorUpdateIndividualInformationCmd defines 'update-individual-information' subcommand
var OperatorUpdateIndividualInformationCmd = &cobra.Command{
	Use:   "update-individual-information",
	Short: TRAPI("/operators/{operator_id}/individual_information:put:summary"),
	Long:  TRAPI(`/operators/{operator_id}/individual_information:put:description`),
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

		param, err := collectOperatorUpdateIndividualInformationCmdParams(ac)
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

func collectOperatorUpdateIndividualInformationCmdParams(ac *apiClient) (*apiParams, error) {
	if OperatorUpdateIndividualInformationCmdOperatorId == "" {
		OperatorUpdateIndividualInformationCmdOperatorId = ac.OperatorID
	}

	body, err := buildBodyForOperatorUpdateIndividualInformationCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if OperatorUpdateIndividualInformationCmdFullName == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "full-name")
		}

	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForOperatorUpdateIndividualInformationCmd("/operators/{operator_id}/individual_information"),
		query:       buildQueryForOperatorUpdateIndividualInformationCmd(),
		contentType: contentType,
		body:        body,
	}, nil
}

func buildPathForOperatorUpdateIndividualInformationCmd(path string) string {

	escapedOperatorId := url.PathEscape(OperatorUpdateIndividualInformationCmdOperatorId)

	path = strReplace(path, "{"+"operator_id"+"}", escapedOperatorId, -1)

	return path
}

func buildQueryForOperatorUpdateIndividualInformationCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForOperatorUpdateIndividualInformationCmd() (string, error) {
	var result map[string]interface{}

	if OperatorUpdateIndividualInformationCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(OperatorUpdateIndividualInformationCmdBody, "@") {
			fname := strings.TrimPrefix(OperatorUpdateIndividualInformationCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if OperatorUpdateIndividualInformationCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(OperatorUpdateIndividualInformationCmdBody)
		}

		if err != nil {
			return "", err
		}

		err = json.Unmarshal(b, &result)
		if err != nil {
			return "", err
		}
	}

	if result == nil {
		result = make(map[string]interface{})
	}

	if OperatorUpdateIndividualInformationCmdAddressLine1 != "" {
		result["addressLine1"] = OperatorUpdateIndividualInformationCmdAddressLine1
	}

	if OperatorUpdateIndividualInformationCmdAddressLine2 != "" {
		result["addressLine2"] = OperatorUpdateIndividualInformationCmdAddressLine2
	}

	if OperatorUpdateIndividualInformationCmdBuilding != "" {
		result["building"] = OperatorUpdateIndividualInformationCmdBuilding
	}

	if OperatorUpdateIndividualInformationCmdCity != "" {
		result["city"] = OperatorUpdateIndividualInformationCmdCity
	}

	if OperatorUpdateIndividualInformationCmdCountryCode != "" {
		result["countryCode"] = OperatorUpdateIndividualInformationCmdCountryCode
	}

	if OperatorUpdateIndividualInformationCmdFullName != "" {
		result["fullName"] = OperatorUpdateIndividualInformationCmdFullName
	}

	if OperatorUpdateIndividualInformationCmdPhoneNumber != "" {
		result["phoneNumber"] = OperatorUpdateIndividualInformationCmdPhoneNumber
	}

	if OperatorUpdateIndividualInformationCmdState != "" {
		result["state"] = OperatorUpdateIndividualInformationCmdState
	}

	if OperatorUpdateIndividualInformationCmdZipCode != "" {
		result["zipCode"] = OperatorUpdateIndividualInformationCmdZipCode
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
