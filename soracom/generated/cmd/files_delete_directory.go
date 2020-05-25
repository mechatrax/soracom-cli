// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"fmt"

	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// FilesDeleteDirectoryCmdPath holds value of 'path' option
var FilesDeleteDirectoryCmdPath string

// FilesDeleteDirectoryCmdScope holds value of 'scope' option
var FilesDeleteDirectoryCmdScope string

func init() {
	FilesDeleteDirectoryCmd.Flags().StringVar(&FilesDeleteDirectoryCmdPath, "path", "", TRAPI("Target directory path"))

	FilesDeleteDirectoryCmd.Flags().StringVar(&FilesDeleteDirectoryCmdScope, "scope", "private", TRAPI("Scope of the request"))
	FilesCmd.AddCommand(FilesDeleteDirectoryCmd)
}

// FilesDeleteDirectoryCmd defines 'delete-directory' subcommand
var FilesDeleteDirectoryCmd = &cobra.Command{
	Use:   "delete-directory",
	Short: TRAPI("/files/{scope}/{path}/:delete:summary"),
	Long:  TRAPI(`/files/{scope}/{path}/:delete:description`),
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

		param, err := collectFilesDeleteDirectoryCmdParams(ac)
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

func collectFilesDeleteDirectoryCmdParams(ac *apiClient) (*apiParams, error) {
	if FilesDeleteDirectoryCmdPath == "" {
		return nil, fmt.Errorf("required parameter '%s' is not specified", "path")
	}

	return &apiParams{
		method: "DELETE",
		path:   buildPathForFilesDeleteDirectoryCmd("/files/{scope}/{path}/"),
		query:  buildQueryForFilesDeleteDirectoryCmd(),
	}, nil
}

func buildPathForFilesDeleteDirectoryCmd(path string) string {

	escapedPath := harvestFilesPathEscape(FilesDeleteDirectoryCmdPath)

	path = strReplace(path, "{"+"path"+"}", escapedPath, -1)

	escapedScope := url.PathEscape(FilesDeleteDirectoryCmdScope)

	path = strReplace(path, "{"+"scope"+"}", escapedScope, -1)

	return path
}

func buildQueryForFilesDeleteDirectoryCmd() url.Values {
	result := url.Values{}

	return result
}
