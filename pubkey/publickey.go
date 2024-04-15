package dapppubkey

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"sourcecode.social/reiver/go-erorr"

	"github.com/reiver/go-dapp/address"
	"github.com/reiver/go-dapp/digest"
	"github.com/reiver/go-dapp/message"
	"github.com/reiver/go-dapp/signature"
)

const (
	errNothing                            = erorr.Error("dapp: no public-key")
	errPublicKeyHexadecimalStringTooShort = erorr.Error("dapp: public-key hexadecimal-string too short")
	errPublicKeyBytesTooShort             = erorr.Error("dapp: public-key bytes too short")
)

type PublicKey struct {
	data []byte
	something bool
}

func NoPublicKey() PublicKey {
	return PublicKey{}
}

func LoadPublicKeyFromBytes(data []byte) (PublicKey, error) {
	return PublicKey{
		something:true,
		data:data,
	}, nil
}

func LoadPublicKeyFromHexadecimalString(hexstr string) (PublicKey, error) {
	var length int = len(hexstr)

	if length < 2 {
		return NoPublicKey(), errPublicKeyHexadecimalStringTooShort
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
			return NoPublicKey(), erorr.Errorf("dapp: problem decoding hexadecimal-string: %w", err)
		}
	}

	return LoadPublicKeyFromBytes(data)
}

func LoadPublicKeyFromMessageAndSignature(message dappmessage.Message, signature dappsignature.Signature) (PublicKey, error) {
	return LoadPublicKeyFromEthereumTextHashDigestAndSignature(message.EthereumTextHashDigest(), signature)
}

func LoadPublicKeyFromEthereumTextHashDigestAndSignature(ethereumTextHashDigest dappdigest.Digest, signature dappsignature.Signature) (PublicKey, error) {

	var signatureBytes []byte = signature.Bytes()
	if 27 == signatureBytes[ethcrypto.RecoveryIDOffset] || 28 == signatureBytes[ethcrypto.RecoveryIDOffset] {
		signatureBytes[ethcrypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	pubKeyData, err := ethcrypto.Ecrecover(ethereumTextHashDigest.Bytes(), signatureBytes)
	if nil != err {
		return NoPublicKey(), erorr.Errorf("dapp: problem with loading pub-key from message and signature: %w", err)
	}

	return LoadPublicKeyFromBytes(pubKeyData)
}

func (receiver PublicKey) Address() (dappaddress.Address, error) {
	if !receiver.something {
		return dappaddress.NoAddress(), errNothing
	}

	var ecdsaPublicKey ecdsa.PublicKey
	{
		var err error

		ecdsaPublicKey, err = receiver.ECDSAPublicKey()
		if nil != err {
			return dappaddress.NoAddress(), err
		}
	}

	var addressFromECDSAPublicKey ethcommon.Address = ethcrypto.PubkeyToAddress(ecdsaPublicKey)

	return dappaddress.LoadAddressFromBytes(addressFromECDSAPublicKey[:])
}

func (receiver PublicKey) Bytes() []byte {
	return append([]byte(nil), receiver.data...)
}

func (receiver PublicKey) ECDSAPublicKey() (ecdsa.PublicKey, error) {

	if len(receiver.data) <= 0 {
		return ecdsa.PublicKey{}, errPublicKeyBytesTooShort
	}

//@TODO: should I be checking if this actually contains the prefix I think it contains.
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
