package emojihash

// package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	"github.com/status-im/status-go/protocol/identity"
	"github.com/status-im/status-go/static"
)

// func main() {
// 	result, _ := GenerateFor("0x04e25da6994ea2dc4ac70727e07eca153ae92bf7609db7befb7ebdceaad348f4fc55bbe90abf9501176301db5aa103fc0eb3bc3750272a26c424a10887db2a7ea8")
// 	// fmt.Println("o, world!")
// 	_ = result
// }

const (
	emojiAlphabetLen = 2757 // 20bytes of data described by 14 emojis requires at least 2757 length alphabet
	// fixme?: 12
	emojiHashLen = 14
)

var emojisAlphabet []string

func GenerateFor(pubkey string) ([]string, error) {
	if len(emojisAlphabet) == 0 {
		alphabet, err := loadAlphabet()
		if err != nil {
			return nil, err
		}
		emojisAlphabet = *alphabet
	}

	// todo?: use SerializePublicKey
	compressedKey, err := identity.ToCompressedKey(pubkey)
	if err != nil {
		return nil, err
	}

	// todo?: String.slice
	// fixme?: slices all even if only slices[1] will be used
	slices, err := identity.Slices(compressedKey)
	if err != nil {
		return nil, err
	}

	slice1 := hex.EncodeToString(slices[1])
	_ = slice1

	return toEmojiHash(new(big.Int).SetBytes(slices[1]), emojiHashLen, &emojisAlphabet) // convert slice to number
}

func loadAlphabet() (*[]string, error) {
	data, err := static.Asset("emojis.txt")
	if err != nil {
		return nil, err
	}

	alphabet := make([]string, 0, emojiAlphabetLen)

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		alphabet = append(alphabet, strings.Replace(scanner.Text(), "\n", "", -1))
	}

	// current alphabet contains more emojis than needed, just in case some emojis needs to be removed
	// make sure only necessary part is loaded
	if len(alphabet) > emojiAlphabetLen {
		alphabet = alphabet[:emojiAlphabetLen]
	}

	return &alphabet, nil
}

func toEmojiHash(value *big.Int, hashLen int, alphabet *[]string) (hash []string, err error) {
	valueBitLen := value.BitLen()
	alphabetLen := new(big.Int).SetInt64(int64(len(*alphabet)))

	indexes := identity.ToBigBase(value, alphabetLen.Uint64())
	if hashLen == 0 {
		hashLen = len(indexes)
	} else if hashLen > len(indexes) {
		prependLen := hashLen - len(indexes)
		for i := 0; i < prependLen; i++ {
			indexes = append([](uint64){0}, indexes...)
		}
	}

	// alphabetLen^hashLen
	possibleCombinations := new(big.Int).Exp(alphabetLen, new(big.Int).SetInt64(int64(hashLen)), nil)

	// 2^valueBitLen
	requiredCombinations := new(big.Int).Exp(new(big.Int).SetInt64(2), new(big.Int).SetInt64(int64(valueBitLen)), nil)

	if possibleCombinations.Cmp(requiredCombinations) == -1 {
		return nil, errors.New("alphabet or hash length is too short to encode given value")
	}

	for _, v := range indexes {
		hash = append(hash, (*alphabet)[v])
	}

	return hash, nil
}
