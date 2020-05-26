// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// FilesListCmdLimit holds value of 'limit' option
var FilesListCmdLimit string

// FilesListCmdPath holds value of 'path' option
var FilesListCmdPath string

// FilesListCmdScope holds value of 'scope' option
var FilesListCmdScope string

func init() {
	FilesListCmd.Flags().StringVar(&FilesListCmdLimit, "limit", "", TRAPI("Num of entries"))

	FilesListCmd.Flags().StringVar(&FilesListCmdPath, "path", "/", TRAPI("Target path"))

	FilesListCmd.Flags().StringVar(&FilesListCmdScope, "scope", "private", TRAPI("Scope of the request"))
	FilesCmd.AddCommand(FilesListCmd)
}

// FilesListCmd defines 'list' subcommand
var FilesListCmd = &cobra.Command{
	Use:   "list",
	Short: TRAPI("/files/{scope}/{path}/:get:summary"),
	Long:  TRAPI(`/files/{scope}/{path}/:get:description`),
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

		param, err := collectFilesListCmdParams(ac)
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

func collectFilesListCmdParams(ac *apiClient) (*apiParams, error) {

	return &apiParams{
		method: "GET",
		path:   buildPathForFilesListCmd("/files/{scope}/{path}/"),
		query:  buildQueryForFilesListCmd(),
	}, nil
}

func buildPathForFilesListCmd(path string) string {

	escapedPath := harvestFilesPathEscape(FilesListCmdPath)

	path = strReplace(path, "{"+"path"+"}", escapedPath, -1)

	escapedScope := url.PathEscape(FilesListCmdScope)

	path = strReplace(path, "{"+"scope"+"}", escapedScope, -1)

	return path
}

func buildQueryForFilesListCmd() url.Values {
	result := url.Values{}

	if FilesListCmdLimit != "" {
		result.Add("limit", FilesListCmdLimit)
	}

	return result
}
