package cmd

import (
	"encoding/json"
	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// UsersPermissionsUpdateCmdDescription holds value of 'description' option
var UsersPermissionsUpdateCmdDescription string

// UsersPermissionsUpdateCmdOperatorId holds value of 'operator_id' option
var UsersPermissionsUpdateCmdOperatorId string

// UsersPermissionsUpdateCmdPermission holds value of 'permission' option
var UsersPermissionsUpdateCmdPermission string

// UsersPermissionsUpdateCmdUserName holds value of 'user_name' option
var UsersPermissionsUpdateCmdUserName string

// UsersPermissionsUpdateCmdBody holds contents of request body to be sent
var UsersPermissionsUpdateCmdBody string

func init() {
	UsersPermissionsUpdateCmd.Flags().StringVar(&UsersPermissionsUpdateCmdDescription, "description", "", TRAPI(""))

	UsersPermissionsUpdateCmd.Flags().StringVar(&UsersPermissionsUpdateCmdOperatorId, "operator-id", "", TRAPI("operator_id"))

	UsersPermissionsUpdateCmd.Flags().StringVar(&UsersPermissionsUpdateCmdPermission, "permission", "", TRAPI(""))

	UsersPermissionsUpdateCmd.Flags().StringVar(&UsersPermissionsUpdateCmdUserName, "user-name", "", TRAPI("user_name"))

	UsersPermissionsUpdateCmd.Flags().StringVar(&UsersPermissionsUpdateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	UsersPermissionsCmd.AddCommand(UsersPermissionsUpdateCmd)
}

// UsersPermissionsUpdateCmd defines 'update' subcommand
var UsersPermissionsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: TRAPI("/operators/{operator_id}/users/{user_name}/permission:put:summary"),
	Long:  TRAPI(`/operators/{operator_id}/users/{user_name}/permission:put:description`),
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

		param, err := collectUsersPermissionsUpdateCmdParams(ac)
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

func collectUsersPermissionsUpdateCmdParams(ac *apiClient) (*apiParams, error) {

	if UsersPermissionsUpdateCmdOperatorId == "" {
		UsersPermissionsUpdateCmdOperatorId = ac.OperatorID
	}

	body, err := buildBodyForUsersPermissionsUpdateCmd()
	if err != nil {
		return nil, err
	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForUsersPermissionsUpdateCmd("/operators/{operator_id}/users/{user_name}/permission"),
		query:       buildQueryForUsersPermissionsUpdateCmd(),
		contentType: "application/json",
		body:        body,
	}, nil
}

func buildPathForUsersPermissionsUpdateCmd(path string) string {

	path = strings.Replace(path, "{"+"operator_id"+"}", UsersPermissionsUpdateCmdOperatorId, -1)

	path = strings.Replace(path, "{"+"user_name"+"}", UsersPermissionsUpdateCmdUserName, -1)

	return path
}

func buildQueryForUsersPermissionsUpdateCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForUsersPermissionsUpdateCmd() (string, error) {
	var result map[string]interface{}

	if UsersPermissionsUpdateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(UsersPermissionsUpdateCmdBody, "@") {
			fname := strings.TrimPrefix(UsersPermissionsUpdateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if UsersPermissionsUpdateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(UsersPermissionsUpdateCmdBody)
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

	if UsersPermissionsUpdateCmdDescription != "" {
		result["description"] = UsersPermissionsUpdateCmdDescription
	}

	if UsersPermissionsUpdateCmdPermission != "" {
		result["permission"] = UsersPermissionsUpdateCmdPermission
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
