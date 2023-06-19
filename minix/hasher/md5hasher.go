package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
	"io"
	"strings"
	"sync"
)

var (
	poolCryptoMD = &sync.Pool{New: func() any {
		return md5.New()
	}}
)

// MD5Crypto 对字符串进行 md5 加密
func MD5Crypto(tags ...string) (string, string) {
	field := strings.Join(tags, "_")
	h := poolCryptoMD.Get().(hash.Hash)
	h.Reset()
	defer poolCryptoMD.Put(h)
	_, _ = io.WriteString(h, field)
	rs := hex.EncodeToString(h.Sum(nil))
	h.Reset()
	return tags[0], rs
}
