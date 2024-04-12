package dapppubkey

import (
	"encoding/hex"
	"fmt"
	"strings"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"sourcecode.social/reiver/go-erorr"

	"github.com/reiver/go-dapp/message"
	"github.com/reiver/go-dapp/signature"
)

const (
	errHexadecimalStringPubKeyTooShort = erorr.Error("dapp: hexadecimal-string pub-key too short")
)

type PubKey struct {
	data []byte
}

func LoadPubKeyFromBytes(data []byte) (PubKey, error) {
	return PubKey{
		data:data,
	}, nil
}

func LoadPubKeyFromHexadecimalString(hexstr string) (PubKey, error) {
	var length int = len(hexstr)

	if length < 2 {
		return PubKey{}, errHexadecimalStringPubKeyTooShort
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
			return PubKey{}, erorr.Errorf("dapp: problem decoding hexadecimal-string: %w", err)
		}
	}

	return LoadPubKeyFromBytes(data)
}

func LoadPubKeyFromMessageAndSignature(message dappmessage.Message, signature dappsignature.Signature) (PubKey, error) {

	pubKeyData, err := ethcrypto.Ecrecover(message.EthereumTextHashDigest(), signature.Bytes())
	if nil != err {
		return PubKey{}, erorr.Errorf("dapp: problem with loading pub-key from message and signature: %w", err)
	}

	return LoadPubKeyFromBytes(pubKeyData)
}

func (receiver PubKey) Bytes() []byte {
	return append([]byte(nil), receiver.data...)
}

func (receiver PubKey) HexadecimalString() string {
	return fmt.Sprintf("0x%x", receiver.data)
}
