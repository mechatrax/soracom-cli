// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// VpgDeleteVpcPeeringConnectionCmdPcxId holds value of 'pcx_id' option
var VpgDeleteVpcPeeringConnectionCmdPcxId string

// VpgDeleteVpcPeeringConnectionCmdVpgId holds value of 'vpg_id' option
var VpgDeleteVpcPeeringConnectionCmdVpgId string

func init() {
	VpgDeleteVpcPeeringConnectionCmd.Flags().StringVar(&VpgDeleteVpcPeeringConnectionCmdPcxId, "pcx-id", "", TRAPI("VPC peering connection ID to be deleted."))

	VpgDeleteVpcPeeringConnectionCmd.Flags().StringVar(&VpgDeleteVpcPeeringConnectionCmdVpgId, "vpg-id", "", TRAPI("Target VPG ID."))
	VpgCmd.AddCommand(VpgDeleteVpcPeeringConnectionCmd)
}

// VpgDeleteVpcPeeringConnectionCmd defines 'delete-vpc-peering-connection' subcommand
var VpgDeleteVpcPeeringConnectionCmd = &cobra.Command{
	Use:   "delete-vpc-peering-connection",
	Short: TRAPI("/virtual_private_gateways/{vpg_id}/vpc_peering_connections/{pcx_id}:delete:summary"),
	Long:  TRAPI(`/virtual_private_gateways/{vpg_id}/vpc_peering_connections/{pcx_id}:delete:description`),
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

		param, err := collectVpgDeleteVpcPeeringConnectionCmdParams(ac)
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

func collectVpgDeleteVpcPeeringConnectionCmdParams(ac *apiClient) (*apiParams, error) {
	if VpgDeleteVpcPeeringConnectionCmdPcxId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "pcx-id")
	}

	if VpgDeleteVpcPeeringConnectionCmdVpgId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "vpg-id")
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForVpgDeleteVpcPeeringConnectionCmd("/virtual_private_gateways/{vpg_id}/vpc_peering_connections/{pcx_id}"),
		query:  buildQueryForVpgDeleteVpcPeeringConnectionCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForVpgDeleteVpcPeeringConnectionCmd(path string) string {

	escapedPcxId := url.PathEscape(VpgDeleteVpcPeeringConnectionCmdPcxId)

	path = strReplace(path, "{"+"pcx_id"+"}", escapedPcxId, -1)

	escapedVpgId := url.PathEscape(VpgDeleteVpcPeeringConnectionCmdVpgId)

	path = strReplace(path, "{"+"vpg_id"+"}", escapedVpgId, -1)

	return path
}

func buildQueryForVpgDeleteVpcPeeringConnectionCmd() url.Values {
	result := url.Values{}

	return result
}
