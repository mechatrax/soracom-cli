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

// EventHandlersUpdateCmdHandlerId holds value of 'handler_id' option
var EventHandlersUpdateCmdHandlerId string

// EventHandlersUpdateCmdBody holds contents of request body to be sent
var EventHandlersUpdateCmdBody string

func init() {
	EventHandlersUpdateCmd.Flags().StringVar(&EventHandlersUpdateCmdHandlerId, "handler-id", "", TRAPI("handler ID"))

	EventHandlersUpdateCmd.Flags().StringVar(&EventHandlersUpdateCmdBody, "body", "", TRCLI("cli.common_params.body.short_help"))
	EventHandlersCmd.AddCommand(EventHandlersUpdateCmd)
}

// EventHandlersUpdateCmd defines 'update' subcommand
var EventHandlersUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: TRAPI("/event_handlers/{handler_id}:put:summary"),
	Long:  TRAPI(`/event_handlers/{handler_id}:put:description`),
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

		param, err := collectEventHandlersUpdateCmdParams(ac)
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

func collectEventHandlersUpdateCmdParams(ac *apiClient) (*apiParams, error) {
	body, err := buildBodyForEventHandlersUpdateCmd()
	if err != nil {
		return nil, err
	}
	contentType := "application/json"

	if EventHandlersUpdateCmdHandlerId == "" {
		if body == "" {

			return nil, fmt.Errorf("required parameter '%s' is not specified", "handler-id")
		}

	}

	return &apiParams{
		method:      "PUT",
		path:        buildPathForEventHandlersUpdateCmd("/event_handlers/{handler_id}"),
		query:       buildQueryForEventHandlersUpdateCmd(),
		contentType: contentType,
		body:        body,
	}, nil
}

func buildPathForEventHandlersUpdateCmd(path string) string {

	escapedHandlerId := url.PathEscape(EventHandlersUpdateCmdHandlerId)

	path = strReplace(path, "{"+"handler_id"+"}", escapedHandlerId, -1)

	return path
}

func buildQueryForEventHandlersUpdateCmd() url.Values {
	result := url.Values{}

	return result
}

func buildBodyForEventHandlersUpdateCmd() (string, error) {
	var result map[string]interface{}

	if EventHandlersUpdateCmdBody != "" {
		var b []byte
		var err error

		if strings.HasPrefix(EventHandlersUpdateCmdBody, "@") {
			fname := strings.TrimPrefix(EventHandlersUpdateCmdBody, "@")
			// #nosec
			b, err = ioutil.ReadFile(fname)
		} else if EventHandlersUpdateCmdBody == "-" {
			b, err = ioutil.ReadAll(os.Stdin)
		} else {
			b = []byte(EventHandlersUpdateCmdBody)
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
