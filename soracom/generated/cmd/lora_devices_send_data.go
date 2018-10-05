package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// LoraDevicesSendDataCmdData holds value of 'data' option
var LoraDevicesSendDataCmdData string

// LoraDevicesSendDataCmdDeviceId holds value of 'device_id' option
var LoraDevicesSendDataCmdDeviceId string

// LoraDevicesSendDataCmdFPort holds value of 'fPort' option
var LoraDevicesSendDataCmdFPort int64

// LoraDevicesSendDataCmdBody holds contents of request body to be sent
var LoraDevicesSendDataCmdBody string

func init() {
	LoraDevicesSendDataCmd.Flags().StringVar(&LoraDevicesSendDataCmdData, "data", "", TRAPI(""))

	LoraDevicesSendDataCmd.Flags().StringVar(&LoraDevicesSendDataCmdDeviceId, "device-id", "", TRAPI("ID of the recipient device."))

	LoraDevicesSendDataCmd.Flags().Int64Var(&LoraDevicesSendDataCmdFPort, "f-port", 0, TRAPI(""))

	LoraDevicesSendDataCmd.Flags().StringVar(&LoraDevicesSendDataCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	LoraDevicesCmd.AddCommand(LoraDevicesSendDataCmd)
}

// LoraDevicesSendDataCmd defines 'send-data' subcommand
var LoraDevicesSendDataCmd = &cobra.Command{
	Use:   "send-data",
	Short: TRAPI("/lora_devices/{device_id}/data:post:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/data:post:description`),
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

		param, err := collectLoraDevicesSendDataCmdParams(ac)
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

func collectLoraDevicesSendDataCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForLoraDevicesSendDataCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "POST",
		path:        buildPathForLoraDevicesSendDataCmd("/lora_devices/{device_id}/data"),
		query:       buildQueryForLoraDevicesSendDataCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForLoraDevicesSendDataCmd(path string) string {

	path = strings.Replace(path, "{"+"device_id"+"}", LoraDevicesSendDataCmdDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesSendDataCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForLoraDevicesSendDataCmd() (string, error) {
	var result map[string]interface{}

	if LoraDevicesSendDataCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LoraDevicesSendDataCmdBody, "@") {
			fname := strings.TrimPrefix(LoraDevicesSendDataCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LoraDevicesSendDataCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LoraDevicesSendDataCmdBody)
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

	if LoraDevicesSendDataCmdData != "" {
		result["data"] = LoraDevicesSendDataCmdData
	}

	if LoraDevicesSendDataCmdFPort != 0 {
		result["fPort"] = LoraDevicesSendDataCmdFPort
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
