package util

import (
	"crypto"
	"encoding/hex"
	"io"
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
	algoMap["MD5SHA1"] = crypto.MD5SHA1
	algoMap["RIPEMD160"] = crypto.RIPEMD160
	algoMap["SHA3_224"] = crypto.SHA3_224
	algoMap["SHA3_256"] = crypto.SHA3_256
	algoMap["SHA3_384"] = crypto.SHA3_384
	algoMap["SHA3_512"] = crypto.SHA3_512
	algoMap["SHA512_224"] = crypto.SHA512_224
	algoMap["SHA512_256"] = crypto.SHA512_256
	algoMap["BLAKE2s_256"] = crypto.BLAKE2s_256
	algoMap["BLAKE2b_256"] = crypto.BLAKE2b_256
	algoMap["BLAKE2b_384"] = crypto.BLAKE2b_384
	algoMap["BLAKE2b_512"] = crypto.BLAKE2b_512
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

func HashStream(algo string, rc io.ReadCloser) string {
	c, ok := algoMap[algo]
	if !ok {
		return ""
	}
	h := c.New()
	io.Copy(h, rc)
	return hex.EncodeToString(h.Sum([]byte{}))
}
