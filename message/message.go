package dappmessage

import (
	"encoding/hex"
	"fmt"
	"strings"

	ethaccounts "github.com/ethereum/go-ethereum/accounts"
	"sourcecode.social/reiver/go-erorr"

	"github.com/reiver/go-dapp/digest"
)

const (
	errHexadecimalStringMessageTooShort = erorr.Error("dapp: hexadecimal-string message too short")
)

type Message struct {
	data []byte
}


func LoadMessageFromBytes(data []byte) (Message, error) {
	return Message{
		data:data,
	}, nil
}

func LoadMessageFromHexadecimalString(hexstr string) (Message, error) {
	var length int = len(hexstr)

	if length < 2 {
		return Message{}, errHexadecimalStringMessageTooShort
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
			return Message{}, erorr.Errorf("dapp: problem decoding hexadecimal-string: %w", err)
		}
	}

	return LoadMessageFromBytes(data)
}

func (receiver Message) Bytes() []byte {
	return append([]byte(nil), receiver.data...)
}

func (receiver Message) HexadecimalString() string {
	return fmt.Sprintf("0x%x", receiver.data)
}

func (receiver Message) EthereumTextHashDigest() dappdigest.Digest {
	var digestBytes []byte = ethaccounts.TextHash(receiver.data)

	digest, _ := dappdigest.LoadDigestFromBytes(digestBytes)

	return digest
}
