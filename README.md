# go-dapp

Package **dapp** provides to for making **dapps**, for the Go programming language.

These include ethereum and other evm based networks.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-dapp

[![GoDoc](https://godoc.org/github.com/reiver/go-dapp?status.svg)](https://godoc.org/github.com/reiver/go-dapp)

## Examples

Here is an example of verifying a signature:
```go
import "github.com/reiver/go-dapp"

// ...

var addressHexadecimalString   string = "--PUT-THE-HEXADECIMAL-ADDRESS-HERE--"
var messageHexadecimalString   string = "--PUT-THE-HEXADECIMAL-MESSAGE-HERE--"
var signatureHexadecimalString string = "--PUT-THE-HEXIDECIMAL-SIGNATURE-HERE--"

// This is the adddress you expect.
address, err := dapp.LoadAddressFromHexadecimalString(addressHexadecimalString)
if nil != err {
	return err
}

message, err := dapp.LoadMessageFromHexadecimalString(messageHexadecimalString)
if nil != err {
	return err
}

signature, err := dapp.LoadSignatureFromHexadecimalString(signatureHexadecimalString)
if nil != err {
	return err
}

addressFromSignature, err := dapp.LoadAddressFromMessageAndSignature(message, signature)
if nil != err {
	return err
}

var verified bool = (address == addressFromSignature)
```

## Import

To import package **dapp** use `import` code like the follownig:
```
import "github.com/reiver/go-dapp"
```

## Installation

To install package **dapp** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-dapp
```

## Author

Package **dapp** was written by [Charles Iliya Krempeaux](http://reiver.link)
