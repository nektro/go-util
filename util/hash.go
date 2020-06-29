package util

import (
	"crypto"
	"encoding/hex"
	"hash"
	"io"

	// ensure crypto algorithms are initialized
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"

	// ensure crypto algorithms are initialized
	_ "golang.org/x/crypto/blake2b"
	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
	_ "golang.org/x/crypto/sha3"
)

var (
	algoMap = map[string]func() hash.Hash{}
)

func init() {
	algoMap["MD4"] = crypto.MD4.New
	algoMap["MD5"] = crypto.MD5.New
	algoMap["SHA1"] = crypto.SHA1.New
	algoMap["SHA224"] = crypto.SHA224.New
	algoMap["SHA256"] = crypto.SHA256.New
	algoMap["SHA384"] = crypto.SHA384.New
	algoMap["SHA512"] = crypto.SHA512.New
	algoMap["MD5SHA1"] = crypto.MD5SHA1.New
	algoMap["RIPEMD160"] = crypto.RIPEMD160.New
	algoMap["SHA3_224"] = crypto.SHA3_224.New
	algoMap["SHA3_256"] = crypto.SHA3_256.New
	algoMap["SHA3_384"] = crypto.SHA3_384.New
	algoMap["SHA3_512"] = crypto.SHA3_512.New
	algoMap["SHA512_224"] = crypto.SHA512_224.New
	algoMap["SHA512_256"] = crypto.SHA512_256.New
	algoMap["BLAKE2s_256"] = crypto.BLAKE2s_256.New
	algoMap["BLAKE2b_256"] = crypto.BLAKE2b_256.New
	algoMap["BLAKE2b_384"] = crypto.BLAKE2b_384.New
	algoMap["BLAKE2b_512"] = crypto.BLAKE2b_512.New
}

func Hash(algo string, bys []byte) string {
	c, ok := algoMap[algo]
	if !ok {
		return ""
	}
	h := c()
	h.Write(bys)
	return hex.EncodeToString(h.Sum([]byte{}))
}

func HashStream(algo string, rc io.ReadCloser) string {
	c, ok := algoMap[algo]
	if !ok {
		return ""
	}
	h := c()
	io.Copy(h, rc)
	return hex.EncodeToString(h.Sum([]byte{}))
}
