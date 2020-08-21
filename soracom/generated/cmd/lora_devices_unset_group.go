// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// LoraDevicesUnsetGroupCmdDeviceId holds value of 'device_id' option
var LoraDevicesUnsetGroupCmdDeviceId string

func init() {
	LoraDevicesUnsetGroupCmd.Flags().StringVar(&LoraDevicesUnsetGroupCmdDeviceId, "device-id", "", TRAPI("Device ID of the target LoRa device."))
	LoraDevicesCmd.AddCommand(LoraDevicesUnsetGroupCmd)
}

// LoraDevicesUnsetGroupCmd defines 'unset-group' subcommand
var LoraDevicesUnsetGroupCmd = &cobra.Command{
	Use:   "unset-group",
	Short: TRAPI("/lora_devices/{device_id}/unset_group:post:summary"),
	Long:  TRAPI(`/lora_devices/{device_id}/unset_group:post:description`),
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

		param, err := collectLoraDevicesUnsetGroupCmdParams(ac)
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

func collectLoraDevicesUnsetGroupCmdParams(ac *apiClient) (*apiParams, error) {
	if LoraDevicesUnsetGroupCmdDeviceId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "device-id")
	}

	return &apiParams{
		method: "POST",
		path:   buildPathForLoraDevicesUnsetGroupCmd("/lora_devices/{device_id}/unset_group"),
		query:  buildQueryForLoraDevicesUnsetGroupCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLoraDevicesUnsetGroupCmd(path string) string {

	escapedDeviceId := url.PathEscape(LoraDevicesUnsetGroupCmdDeviceId)

	path = strReplace(path, "{"+"device_id"+"}", escapedDeviceId, -1)

	return path
}

func buildQueryForLoraDevicesUnsetGroupCmd() url.Values {
	result := url.Values{}

	return result
}
