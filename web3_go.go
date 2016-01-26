package main

import (
    "fmt"
    "github.com/parnurzeal/gorequest"
    "github.com/bitly/go-simplejson"
    "log"
)

const (
  apiUrl string = "http://localhost:8548" // 8545 (default port)
)

func main() {
  account := coinbase()
  fmt.Println("account:", account)

  accs := accounts()
  fmt.Println("accounts:", accs)

  bal := getBalance(account)
  fmt.Println("balance of account:", account, "=", bal)

  // TODO:

  // eth_call

  // eth_sendTransaction - https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sendtransaction

  // eth_getTransactionReceipt

  // eth_sign?
}

func coinbase() (string) {
  res := call("eth_coinbase").Get("result").MustString()
  return res
}

func accounts() ([]interface {}) {
  res := call("eth_accounts").Get("result").MustArray()
  return res
}

func getBalance(address string) (int) {
  res := call("eth_getBalance", "["+address+"]").Get("result").MustInt()
  return res
}

func call(args ...string) (*simplejson.Json) {
  method := args[0]
  params := "[]"
  if len(args) > 1 {
    params = args[1]
  }

  _, body, errs := gorequest.New().Post(apiUrl).
    Send(`{"jsonrpc":"2.0","method":"`+method+`","params":`+params+`}`).
    End()

  if errs != nil {
    panic(errs)
  }

  js, err := simplejson.NewJson([]byte(body))
  if err != nil {
      log.Fatalln(err)
  }

  return js
}
