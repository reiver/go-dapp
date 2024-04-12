package dappsignature

import (
	"encoding/hex"
	"fmt"
	"strings"

	"sourcecode.social/reiver/go-erorr"
)

const (
	errHexadecimalStringSignatureTooShort = erorr.Error("dapp: hexadecimal-signature too short")
)

type Signature struct {
	data []byte
}

func LoadSignatureFromBytes(p []byte) (Signature, error) {
	var data []byte = append([]byte(nil), p...)

	return Signature{
		data:data,
	}, nil
}

func LoadSignatureFromHexadecimal(hexstr string) (Signature, error) {
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

	return Signature{
		data:data,
	}, nil
}

func (receiver Signature) Bytes() []byte {
	return append([]byte(nil), receiver.data...)
}

func (receiver Signature) String() string {
	return fmt.Sprintf("0x%x", receiver.data)
}
