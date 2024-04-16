package dapp

import (
	"encoding/hex"
	"fmt"
	"strings"

	"sourcecode.social/reiver/go-erorr"
)

const (
	errHexadecimalStringDigestTooShort = erorr.Error("dapp: hexadecimal-string digest too short")
)

type Digest struct {
	data []byte
}


func LoadDigestFromBytes(data []byte) (Digest, error) {
	return Digest{
		data:data,
	}, nil
}

func LoadDigestFromHexadecimalString(hexstr string) (Digest, error) {
	var length int = len(hexstr)

	if length < 2 {
		return Digest{}, errHexadecimalStringDigestTooShort
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
			return Digest{}, erorr.Errorf("dapp: problem decoding hexadecimal-string: %w", err)
		}
	}

	return LoadDigestFromBytes(data)
}

func (receiver Digest) Bytes() []byte {
	return append([]byte(nil), receiver.data...)
}

func (receiver Digest) HexadecimalString() string {
	return fmt.Sprintf("0x%x", receiver.data)
}
