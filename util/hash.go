package util

import (
	"crypto"
	"encoding/hex"
)

var (
	algoMap = map[string]crypto.Hash{}
)

func init() {
	algoMap["MD4"] = crypto.MD4
	algoMap["MD5"] = crypto.MD5
	algoMap["SHA1"] = crypto.SHA1
	algoMap["SHA224"] = crypto.SHA224
	algoMap["SHA256"] = crypto.SHA256
	algoMap["SHA384"] = crypto.SHA384
	algoMap["SHA512"] = crypto.SHA512
	algoMap["RIPEMD160"] = crypto.RIPEMD160
}

func Hash(algo string, bys []byte) string {
	c, ok := algoMap[algo]
	if !ok {
		return ""
	}
	h := c.New()
	h.Write(bys)
	return hex.EncodeToString(h.Sum([]byte{}))
}
