package main

import (
    "fmt"
    "./web3"
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
  res := web3.Call("eth_coinbase").Get("result").MustString()
  return res
}

func accounts() ([]interface {}) {
  res := web3.Call("eth_accounts").Get("result").MustArray()
  return res
}

func getBalance(address string) (int) {
  res := web3.Call("eth_getBalance", "["+address+"]").Get("result").MustInt()
  return res
}
