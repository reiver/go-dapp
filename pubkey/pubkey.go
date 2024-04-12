package dapppubkey

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"sourcecode.social/reiver/go-erorr"

	"github.com/reiver/go-dapp/digest"
	"github.com/reiver/go-dapp/message"
	"github.com/reiver/go-dapp/signature"
)

const (
	errPublicKeyHexadecimalStringTooShort = erorr.Error("dapp: public-key hexadecimal-string too short")
	errPublicKeyBytesTooShort             = erorr.Error("dapp: public-key bytes too short")
)

type PublicKey struct {
	data []byte
}

func LoadPublicKeyFromBytes(data []byte) (PublicKey, error) {
	return PublicKey{
		data:data,
	}, nil
}

func LoadPublicKeyFromHexadecimalString(hexstr string) (PublicKey, error) {
	var length int = len(hexstr)

	if length < 2 {
		return PublicKey{}, errPublicKeyHexadecimalStringTooShort
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
			return PublicKey{}, erorr.Errorf("dapp: problem decoding hexadecimal-string: %w", err)
		}
	}

	return LoadPublicKeyFromBytes(data)
}

func LoadPublicKeyFromMessageAndSignature(message dappmessage.Message, signature dappsignature.Signature) (PublicKey, error) {
	return LoadPublicKeyFromEthereumTextHashDigestAndSignature(message.EthereumTextHashDigest(), signature)
}

func LoadPublicKeyFromEthereumTextHashDigestAndSignature(ethereumTextHashDigest dappdigest.Digest, signature dappsignature.Signature) (PublicKey, error) {

	pubKeyData, err := ethcrypto.Ecrecover(ethereumTextHashDigest.Bytes(), signature.Bytes())
	if nil != err {
		return PublicKey{}, erorr.Errorf("dapp: problem with loading pub-key from message and signature: %w", err)
	}

	return LoadPublicKeyFromBytes(pubKeyData)
}

func (receiver PublicKey) Bytes() []byte {
	return append([]byte(nil), receiver.data...)
}

func (receiver PublicKey) ECDSAPublicKey() (ecdsa.PublicKey, error) {

	if len(receiver.data) <= 0 {
		return ecdsa.PublicKey{}, errPublicKeyBytesTooShort
	}

//@TODO: should I be checking if this actually contrains the prefix I think it contains.
	// Remove the prefix (0x04 for uncompressed key)
	var truncatedPublicKeyBytes []byte = receiver.data[1:]

	var ecdsaPublicKey ecdsa.PublicKey
	ecdsaPublicKey.X = new(big.Int).SetBytes(truncatedPublicKeyBytes[:32])
	ecdsaPublicKey.Y = new(big.Int).SetBytes(truncatedPublicKeyBytes[32:])

	return ecdsaPublicKey, nil
}

func (receiver PublicKey) HexadecimalString() string {
	return fmt.Sprintf("0x%x", receiver.data)
}
