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
	prefix    string
}

func NewHasher(key string, prefix string) *Hasher {
	return &Hasher{
		key:       []byte(key),
		md5Hasher: md5.New(),
		prefix:    prefix,
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
	prefixLength := len(hasher.prefix)

	if len(result) > prefixLength+1 {
		if result[0:prefixLength] == hasher.prefix {
			return true
		}
	}

	return false
}
