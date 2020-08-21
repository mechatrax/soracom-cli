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

// LagoonDashboardsUpdatePermissionsCmdDashboardId holds value of 'dashboard_id' option
var LagoonDashboardsUpdatePermissionsCmdDashboardId int64

// LagoonDashboardsUpdatePermissionsCmdBody holds contents of request body to be sent
var LagoonDashboardsUpdatePermissionsCmdBody string

func init() {
	LagoonDashboardsUpdatePermissionsCmd.Flags().Int64Var(&LagoonDashboardsUpdatePermissionsCmdDashboardId, "dashboard-id", 0, TRAPI("dashboard_id"))

	LagoonDashboardsUpdatePermissionsCmd.Flags().StringVar(&LagoonDashboardsUpdatePermissionsCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	LagoonDashboardsCmd.AddCommand(LagoonDashboardsUpdatePermissionsCmd)
}

// LagoonDashboardsUpdatePermissionsCmd defines 'update-permissions' subcommand
var LagoonDashboardsUpdatePermissionsCmd = &cobra.Command{
	Use:   "update-permissions",
	Short: TRAPI("/lagoon/dashboards/{dashboard_id}/permissions:put:summary"),
	Long:  TRAPI(`/lagoon/dashboards/{dashboard_id}/permissions:put:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectLagoonDashboardsUpdatePermissionsCmdParams(ac)
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

func collectLagoonDashboardsUpdatePermissionsCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForLagoonDashboardsUpdatePermissionsCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if LagoonDashboardsUpdatePermissionsCmdDashboardId == 0 {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "dashboard-id")
		}

	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForLagoonDashboardsUpdatePermissionsCmd("/lagoon/dashboards/{dashboard_id}/permissions"),
		query:       buildQueryForLagoonDashboardsUpdatePermissionsCmd(),
		contentType: contentType,
		body:        body,

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForLagoonDashboardsUpdatePermissionsCmd(path string) string {

	path = strReplace(path, "{"+"dashboard_id"+"}", url.PathEscape(sprintf("%d", LagoonDashboardsUpdatePermissionsCmdDashboardId)), -1)

	return path
}

func buildQueryForLagoonDashboardsUpdatePermissionsCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForLagoonDashboardsUpdatePermissionsCmd() (string, error) {
	var result map[string]interface{}

	if LagoonDashboardsUpdatePermissionsCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(LagoonDashboardsUpdatePermissionsCmdBody, "@") {
			fname := strings.TrimPrefix(LagoonDashboardsUpdatePermissionsCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if LagoonDashboardsUpdatePermissionsCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(LagoonDashboardsUpdatePermissionsCmdBody)
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

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
