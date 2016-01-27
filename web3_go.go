package main

import (
    "fmt"
    "github.com/makevoid/web3_go/web3"
    "github.com/bitly/go-simplejson"
    "encoding/json"
    "golang.org/x/crypto/sha3"
    "encoding/hex"
)

func main() {
  account := coinbase()
  fmt.Println("account:", account)

  accs := accounts()
  fmt.Println("accounts:", accs)

  bal := getBalance(account)
  fmt.Println("balance of account:", account, "=", bal)

  contract := "contract test { function multiply(uint a) returns(uint d) {   return a * 7;   } }"
  resp_c := compile(contract)
  fmt.Println("compiled contract infos (abi):")
  info_c := getAbi(resp_c)
  pp(info_c)

  resp := call()
  fmt.Println("call method(), resp:", resp)
  pp(resp)

  sha := sha3.Sum512( []byte("foo") )
  fmt.Println("sha3('foo'):", sha)
  shaHex := hex.EncodeToString(sha[:])
  // fmt.Println("sha3('foo'):", string(sha[:32]))
  fmt.Println("sha3('foo'):", shaHex)

  // TODO:
  // eth_sendTransaction - https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sendtransaction
  // eth_getTransactionReceipt
  // eth_sign?

  // 4bca2b137edc580fe50a88983ef860ebaca36c857b1f492839d6d7392452a63c82cbebc68e3b70a2a1480b4bb5d437a7cba6ecf9d89f9ff3ccd14cd6146ea7e7
  // 1597842aac52bc9d13fe249d808afbf44da13524759477404c3592ee331173e89fe1cbf21a7e4360990d565fad4643cdb209d80fa41a91dea97e665022c92135
}

func compile(contract string) (*simplejson.Json) {
  res := web3.Call("eth_compileSolidity", `["`+contract+`"]`).Get("result")
  return res
}

func getAbi(compiledResp *simplejson.Json) (interface {}) {
  return compiledResp.Get("test").Get("info").Get("abiDefinition").MustArray()
}

func call() (interface {}) {
  to := "0x5e7565ff99945b476a830c04ec15bc631362bcd7"
  data := "0x3e27986000000000000000000000000000000002400000000000000000000000000000000000000000000000000000000000000880000000000000000000000000000000"
  res := web3.Call("eth_call", `{ "to": "`+to+`", "data": "`+data+`", }`).Get("result").MustString()
  return res
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

// pretty print

func pp(data interface {}) {
  js, err := json.MarshalIndent(data, "", "  ")
  if err != nil {
      fmt.Println("error:", err)
  }
  fmt.Print(string(js))
  println()
}
