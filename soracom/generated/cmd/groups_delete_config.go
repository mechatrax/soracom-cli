// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// GroupsDeleteConfigCmdGroupId holds value of 'group_id' option
var GroupsDeleteConfigCmdGroupId string

// GroupsDeleteConfigCmdName holds value of 'name' option
var GroupsDeleteConfigCmdName string

// GroupsDeleteConfigCmdNamespace holds value of 'namespace' option
var GroupsDeleteConfigCmdNamespace string

func init() {
	GroupsDeleteConfigCmd.Flags().StringVar(&GroupsDeleteConfigCmdGroupId, "group-id", "", TRAPI("Target group."))

	GroupsDeleteConfigCmd.Flags().StringVar(&GroupsDeleteConfigCmdName, "name", "", TRAPI("Parameter name to be deleted. (This will be part of a URL path, so it needs to be percent-encoded. In JavaScript, specify the name after it has been encoded using encodeURIComponent().)"))

	GroupsDeleteConfigCmd.Flags().StringVar(&GroupsDeleteConfigCmdNamespace, "namespace", "", TRAPI("Namespace of target parameters."))
	GroupsCmd.AddCommand(GroupsDeleteConfigCmd)
}

// GroupsDeleteConfigCmd defines 'delete-config' subcommand
var GroupsDeleteConfigCmd = &cobra.Command{
	Use:   "delete-config",
	Short: TRAPI("/groups/{group_id}/configuration/{namespace}/{name}:delete:summary"),
	Long:  TRAPI(`/groups/{group_id}/configuration/{namespace}/{name}:delete:description`),
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

		param, err := collectGroupsDeleteConfigCmdParams(ac)
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

func collectGroupsDeleteConfigCmdParams(ac *apiClient) (*apiParams, error) {
	if GroupsDeleteConfigCmdGroupId == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "group-id")
	}

	if GroupsDeleteConfigCmdName == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "name")
	}

	if GroupsDeleteConfigCmdNamespace == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "namespace")
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForGroupsDeleteConfigCmd("/groups/{group_id}/configuration/{namespace}/{name}"),
		query:  buildQueryForGroupsDeleteConfigCmd(),

		noRetryOnError: noRetryOnError,
	}, nil
}

func buildPathForGroupsDeleteConfigCmd(path string) string {

	escapedGroupId := url.PathEscape(GroupsDeleteConfigCmdGroupId)

	path = strReplace(path, "{"+"group_id"+"}", escapedGroupId, -1)

	escapedName := url.PathEscape(GroupsDeleteConfigCmdName)

	path = strReplace(path, "{"+"name"+"}", escapedName, -1)

	escapedNamespace := url.PathEscape(GroupsDeleteConfigCmdNamespace)

	path = strReplace(path, "{"+"namespace"+"}", escapedNamespace, -1)

	return path
}

func buildQueryForGroupsDeleteConfigCmd() url.Values {
	result := url.Values{}

	return result
}
