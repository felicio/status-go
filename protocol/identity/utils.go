package identity

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func ToColorID(pubkey string) (int64, error) {
	const colorPalletLength = 12

	pubkeyValue, ok := new(big.Int).SetString(pubkey, 0)
	if !ok {
		return 0, fmt.Errorf("invalid pubkey: %s", pubkey)
	}

	colorID := new(big.Int).Mod(pubkeyValue, new(big.Int).SetInt64(colorPalletLength-1)).Int64()

	return colorID, nil
}

func ToBigBase(value *big.Int, base uint64) (res [](uint64)) {
	toBigBaseImpl(value, base, &res)
	return
}

// normalizing to base +2800|-2757
// (33*8)/11 (24 emojis)
// not the same length of the hash
// changes memory value of the result (pointer)
func toBigBaseImpl(value *big.Int, base uint64, res *[](uint64)) {
	bigBase := new(big.Int).SetUint64(base)
	quotient := new(big.Int).Div(value, bigBase)
	// tmp := new(big.Int).SetUint64(0)
	if quotient.Cmp(new(big.Int).SetUint64(0)) != 0 { // until it's 0 || is less then base || float vs int || value left over is less than base
		toBigBaseImpl(quotient, base, res)
	}

	// whole round number 1-2800
	// mod; give reaminder; useful for indexing
	*res = append(*res, new(big.Int).Mod(value, bigBase).Uint64()) // gives index on the base value
}

// compressedPubKey = |1.5 bytes chars cutoff|20 bytes emoji hash|10 bytes color hash|1.5 bytes chars cutoff|
func Slices(compressedPubkey []byte) (res [4][]byte, err error) {
	// compressedPubkeyInString := hex.EncodeToString((compressedPubkey))
	// _ = compressedPubkeyInString

	if len(compressedPubkey) != 33 {
		return res, errors.New("incorrect compressed pubkey")
	}

	getSlice := func(low, high int, and string, rsh uint) []byte {
		sliceValue := new(big.Int).SetBytes(compressedPubkey[low:high])
		andValue, _ := new(big.Int).SetString(and, 0)
		andRes := new(big.Int).And(sliceValue, andValue)
		return new(big.Int).Rsh(andRes, rsh).Bytes()
	}

	res[0] = getSlice(0, 2, "0xFFF0", 4)
	res[1] = getSlice(1, 22, "0x0FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF0", 4)
	res[2] = getSlice(21, 32, "0x0FFFFFFFFFFFFFFFFFFFF0", 4)
	res[3] = getSlice(31, 33, "0x0FFF", 0)

	return res, nil
}

func ToCompressedKey(pubkey string) ([]byte, error) {
	pubkeyValue, ok := new(big.Int).SetString(pubkey, 0)
	if !ok {
		return nil, fmt.Errorf("invalid pubkey: %s", pubkey)
	}

	pubkeyValueInBytes := pubkeyValue.Bytes()
	x, y := secp256k1.S256().Unmarshal(pubkeyValueInBytes)
	if x == nil || !secp256k1.S256().IsOnCurve(x, y) {
		return nil, fmt.Errorf("invalid pubkey: %s", pubkey)
	}

	result := secp256k1.CompressPubkey(x, y)

	// resultInString := string(result)
	resultInString := hex.EncodeToString(result)
	_ = resultInString // silence "resultInString declared but not used (exit status 2)"

	return result, nil
}

func ToBigInt(t *testing.T, str string) *big.Int {
	res, ok := new(big.Int).SetString(str, 0)
	if !ok {
		t.Errorf("invalid conversion to int from %s", str)
	}
	return res
}
