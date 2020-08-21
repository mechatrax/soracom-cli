// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LoraDevicesEnableTerminationCmdDeviceId holds value of 'device_id' option
var LoraDevicesEnableTerminationCmdDeviceId string

func init() {
	LoraDevicesEnableTerminationCmd.Flags().StringVar(&LoraDevicesEnableTerminationCmdDeviceId, "device-id", "", TRAPI("Device ID of the target LoRa device."))
	LoraDevicesCmd.AddCommand(LoraDevicesEnableTerminationCmd)
}

// LoraDevicesEnableTerminationCmd defines 'enable-termination' subcommand
var LoraDevicesEnableTerminationCmd = &cobra.Command{
	Use:   "enable-termination",
	Short: TRAPI("/lora_devices/{device_id}/enable_termination:post:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/enable_termination:post:description`),
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

		param, err := collectLoraDevicesEnableTerminationCmdParams(ac)
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

func collectLoraDevicesEnableTerminationCmdParams(ac *apiClient) (*apiParams, error) {
	if LoraDevicesEnableTerminationCmdDeviceId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "device-id")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraDevicesEnableTerminationCmd("/lora_devices/{device_id}/enable_termination"),
		query:  buildQueryForLoraDevicesEnableTerminationCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraDevicesEnableTerminationCmd(path string) string {

	escapedDeviceId := url.PathEscape(LoraDevicesEnableTerminationCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesEnableTerminationCmd() url.Values {
	result := url.Values{}

	return result
}
