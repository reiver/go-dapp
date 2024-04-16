package dapp

import (
	"encoding/hex"
	"fmt"
	"strings"

	"sourcecode.social/reiver/go-erorr"
)

const (
	errHexadecimalStringSignatureTooShort = erorr.Error("dapp: hexadecimal-string signature too short")
)

type Signature struct {
	data []byte
}

func LoadSignatureFromBytes(data []byte) (Signature, error) {
	return Signature{
		data:data,
	}, nil
}

func LoadSignatureFromHexadecimalString(hexstr string) (Signature, error) {
	var length int = len(hexstr)

	if length < 2 {
		return Signature{}, errHexadecimalStringSignatureTooShort
	}

	{
		const prefix string = "0x"

		if strings.HasPrefix(hexstr, prefix) {
			hexstr = hexstr[2:]
		}
	}

	var data []byte
	{
		var err error

		data, err = hex.DecodeString(hexstr)
		if nil != err {
			return Signature{}, erorr.Errorf("dapp: problem decoding hexadecimal-string: %w", err)
		}
	}

	return LoadSignatureFromBytes(data)
}

func (receiver Signature) SigningAddress(message Message) (Address, error) {

	publicKeyFromSignature, err := LoadPublicKeyFromMessageAndSignature(message, receiver)
	if nil != err {
		return NoAddress(), err
	}

	return publicKeyFromSignature.Address()
}

func (receiver Signature) Bytes() []byte {
	return append([]byte(nil), receiver.data...)
}

func (receiver Signature) HexadecimalString() string {
	return fmt.Sprintf("0x%x", receiver.data)
}
