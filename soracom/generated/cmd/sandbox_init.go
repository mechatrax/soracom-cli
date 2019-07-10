// Code generated by soracom-cli generate-cmd. DO NOT EDIT.
package cmd

import (
	"encoding/json"

	"io/ioutil"

	"os"
	"strings"

	"github.com/spf13/cobra"
)

// SandboxInitCmdAuthKey holds value of 'authKey' option
var SandboxInitCmdAuthKey string

// SandboxInitCmdAuthKeyId holds value of 'authKeyId' option
var SandboxInitCmdAuthKeyId string

// SandboxInitCmdEmail holds value of 'email' option
var SandboxInitCmdEmail string

// SandboxInitCmdPassword holds value of 'password' option
var SandboxInitCmdPassword string

// SandboxInitCmdBody holds contents of request body to be sent
var SandboxInitCmdBody string

func init() {
	SandboxInitCmd.Flags().StringVar(&SandboxInitCmdAuthKey, "auth-key", "", TRAPI(""))

	SandboxInitCmd.MarkFlagRequired("auth-key")

	SandboxInitCmd.Flags().StringVar(&SandboxInitCmdAuthKeyId, "auth-key-id", "", TRAPI(""))

	SandboxInitCmd.MarkFlagRequired("auth-key-id")

	SandboxInitCmd.Flags().StringVar(&SandboxInitCmdEmail, "email", "", TRAPI(""))

	SandboxInitCmd.MarkFlagRequired("email")

	SandboxInitCmd.Flags().StringVar(&SandboxInitCmdPassword, "password", "", TRAPI(""))

	SandboxInitCmd.MarkFlagRequired("password")

	SandboxInitCmd.Flags().StringVar(&SandboxInitCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))

	SandboxCmd.AddCommand(SandboxInitCmd)
}

// SandboxInitCmd defines 'init' subcommand
var SandboxInitCmd = &cobra.Command{
	Use:   "init",
	Short: TRAPI("/sandbox/init:post:summary"),
	Long:  TRAPI(`/sandbox/init:post:description`),
	RunE: func(cmd *cobra.Command, args []string) error {
		opt := &apiClientOptions{
			BasePath: "/v1",
			Language: getSelectedLanguage(),
		}

		ac := newAPIClient(opt)
		if v := os.Getenv("SORACOM_VERBOSE"); v != "" {
			ac.SetVerbose(true)
		}

		param, err := collectSandboxInitCmdParams(ac)
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

func collectSandboxInitCmdParams(ac *apiClient) (*apiParams, error) {

	body, err := buildBodyForSandboxInitCmd()
	if err != nil {
		return nil, err
	}

	contentType := "application/json"

	return &apiParams{
		method:      "POST",
		path:        buildPathForSandboxInitCmd("/sandbox/init"),
		query:       buildQueryForSandboxInitCmd(),
		contentType: contentType,
		body:        body,
	}, nil
}

func buildPathForSandboxInitCmd(path string) string {

	return path
}

func buildQueryForSandboxInitCmd() string {
	result := []string{}

	return strings.Join(result, "&")
}

func buildBodyForSandboxInitCmd() (string, error) {
	var result map[string]interface{}

	if SandboxInitCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(SandboxInitCmdBody, "@") {
			fname := strings.TrimPrefix(SandboxInitCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if SandboxInitCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(SandboxInitCmdBody)
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

	if SandboxInitCmdAuthKey != "" {
		result["authKey"] = SandboxInitCmdAuthKey
	}

	if SandboxInitCmdAuthKeyId != "" {
		result["authKeyId"] = SandboxInitCmdAuthKeyId
	}

	if SandboxInitCmdEmail != "" {
		result["email"] = SandboxInitCmdEmail
	}

	if SandboxInitCmdPassword != "" {
		result["password"] = SandboxInitCmdPassword
	}

	resultBytes, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(resultBytes), nil
}
