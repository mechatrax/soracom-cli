package cmd

import (

  "encoding/json"

  "fmt"

  "io/ioutil"

  "os"
  "strings"

  "github.com/spf13/cobra"
)






var PaymentMethodsWebpayRegisterCmdCvc string

var PaymentMethodsWebpayRegisterCmdName string

var PaymentMethodsWebpayRegisterCmdNumber string




var PaymentMethodsWebpayRegisterCmdBody string


func init() {
  PaymentMethodsWebpayRegisterCmd.Flags().StringVar(&PaymentMethodsWebpayRegisterCmdCvc, "cvc", "", "")

  PaymentMethodsWebpayRegisterCmd.Flags().StringVar(&PaymentMethodsWebpayRegisterCmdName, "name", "", "")

  PaymentMethodsWebpayRegisterCmd.Flags().StringVar(&PaymentMethodsWebpayRegisterCmdNumber, "number", "", "")



  PaymentMethodsWebpayRegisterCmd.Flags().StringVar(&PaymentMethodsWebpayRegisterCmdBody, "body", "", TR("cli.common_params.body.short_help"))


  PaymentMethodsWebpayCmd.AddCommand(PaymentMethodsWebpayRegisterCmd)
}

var PaymentMethodsWebpayRegisterCmd = &cobra.Command{
  Use: "register",
  Short: TR("Register credit card information for WebPay"),
  Long: TR(`WebPayでの支払い用のカード情報を登録します`),
  RunE: func(cmd *cobra.Command, args []string) error {
    opt := &apiClientOptions{
      Endpoint: getSpecifiedEndpoint(),
      BasePath: "/v1",
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
    
    param, err := collectPaymentMethodsWebpayRegisterCmdParams()
    if err != nil {
      return err
    }

    result, err := ac.callAPI(param)
    if err != nil {
      cmd.SilenceUsage = true
      return err
    }

    fmt.Println(result)
    return nil
  },
}

func collectPaymentMethodsWebpayRegisterCmdParams() (*apiParams, error) {
  
  body, err := buildBodyForPaymentMethodsWebpayRegisterCmd()
  if err != nil {
    return nil, err
  }
  

  return &apiParams{
    method: "POST",
    path: buildPathForPaymentMethodsWebpayRegisterCmd("/payment_methods/webpay"),
    query: buildQueryForPaymentMethodsWebpayRegisterCmd(),
    contentType: "application/json",
    body: body,
  }, nil
}

func buildPathForPaymentMethodsWebpayRegisterCmd(path string) string {
  
  
  
  
  
  
  
  
  
  return path
}

func buildQueryForPaymentMethodsWebpayRegisterCmd() string {
  result := []string{}
  
  
  
  
  
  
  

  

  
  return strings.Join(result, "&")
}


func buildBodyForPaymentMethodsWebpayRegisterCmd() (string, error) {
  if PaymentMethodsWebpayRegisterCmdBody != "" {
    if strings.HasPrefix(PaymentMethodsWebpayRegisterCmdBody, "@") {
      fname := strings.TrimPrefix(PaymentMethodsWebpayRegisterCmdBody, "@")
      bytes, err := ioutil.ReadFile(fname)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else if PaymentMethodsWebpayRegisterCmdBody == "-" {
      bytes, err := ioutil.ReadAll(os.Stdin)
      if err != nil {
        return "", err
      }
      return string(bytes), nil
    } else {
      return PaymentMethodsWebpayRegisterCmdBody, nil
    }
  }

  result := map[string]interface{}{}
  
  if PaymentMethodsWebpayRegisterCmdCvc != "" {
    result["cvc"] = PaymentMethodsWebpayRegisterCmdCvc
  }
  
  if PaymentMethodsWebpayRegisterCmdName != "" {
    result["name"] = PaymentMethodsWebpayRegisterCmdName
  }
  
  if PaymentMethodsWebpayRegisterCmdNumber != "" {
    result["number"] = PaymentMethodsWebpayRegisterCmdNumber
  }
  
  
  

  resultBytes, err := json.Marshal(result)
  if err != nil {
    return "", err
  }
  return string(resultBytes), nil
}

