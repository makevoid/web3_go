package main

import (
    "fmt"
    "github.com/parnurzeal/gorequest"
    "github.com/bitly/go-simplejson"
    "log"
)

const (
  apiUrl string = "http://localhost:8548"
)


func main() {
  account := coinbase()
  fmt.Println(account)

  accs := accounts()
  fmt.Println(accs)
}

func coinbase() (string) {
  body := call("eth_coinbase").Get("result").MustString()
  return body
}

func accounts() ([]interface {}) {
  body2 := call("eth_accounts").Get("result").MustArray()
  return body2
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
