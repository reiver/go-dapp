package dappaddress

import (
	"encoding/hex"
	"fmt"
	"strings"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"sourcecode.social/reiver/go-erorr"
)

const (
	errHexadecimalStringAddressTooShort = erorr.Error("dapp: hexadecimal-string 'address' too short")
)

type Address struct {
	data ethcommon.Address
	something bool
}

func NoAddress() Address {
	return Address{}
}

func LoadAddressFromBytes(data []byte) (Address, error) {

	const addresslength int = ethcommon.AddressLength
	var   datalength    int = len(data)

	if datalength != addresslength {
		return NoAddress(), erorr.Errorf("dapp: bytes for address wrong size: expected length of data for 'address' to be %d bytes but was actually %d bytes", addresslength, datalength)
	}

	var address Address
	address.something = true
	copy(address.data[:], data)

	return address, nil
}

func LoadAddressFromHexadecimalString(hexstr string) (Address, error) {

	var length int = len(hexstr)

	if length < 2 {
		return NoAddress(), errHexadecimalStringAddressTooShort
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
			return NoAddress(), erorr.Errorf("dapp: problem decoding hexadecimal-string: %w", err)
		}
	}

	return LoadAddressFromBytes(data)
}

func (receiver Address) Bytes() []byte {
	return append([]byte(nil), receiver.data[:]...)
}

func (receiver Address) HexadecimalString() string {
	return fmt.Sprintf("0x%x", receiver.data)
}
