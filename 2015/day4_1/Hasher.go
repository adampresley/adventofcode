package main

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
	"strconv"
)

type Hasher struct {
	key       []byte
	md5Hasher hash.Hash
}

func NewHasher(key string) *Hasher {
	return &Hasher{
		key:       []byte(key),
		md5Hasher: md5.New(),
	}
}

func (hasher *Hasher) Compute(value int) string {
	valueInBytes := hasher.intToBytes(value)
	hasher.md5Hasher.Reset()

	hasher.md5Hasher.Write(hasher.key)
	hasher.md5Hasher.Write(valueInBytes)

	return hex.EncodeToString(hasher.md5Hasher.Sum(nil))
}

func (hasher *Hasher) intToBytes(value int) []byte {
	bytes := []byte(strconv.Itoa(value))
	return bytes
}

func (hasher *Hasher) IsValidAnswer(result string) bool {
	if len(result) > 7 {
		if result[0:5] == "00000" {
			sixthDigit, _ := strconv.Atoi(result[5:6])

			if sixthDigit > 0 {
				return true
			}
		}
	}

	return false
}
