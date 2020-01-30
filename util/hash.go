package util

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"

	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
)

var (
	algoMap = map[string]hash.Hash{}
)

func init() {
	algoMap["MD4"] = md4.New()
	algoMap["MD5"] = md5.New()
	algoMap["SHA1"] = sha1.New()
	algoMap["SHA256"] = sha256.New()
	algoMap["SHA512"] = sha512.New()
	algoMap["RIPEMD160"] = ripemd160.New()
}

func Hash(algo string, bys []byte) string {
	c, ok := algoMap[algo]
	if !ok {
		return ""
	}
	defer c.Reset()
	return hex.EncodeToString(c.Sum(bys)[len(bys):])
}
