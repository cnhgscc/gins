package hasher

import (
	"fmt"
	"hash/crc32"
	"strings"
)

func Crc32Crypto(tags ...string) (string, string) {
	field := strings.Join(tags, "_")
	crc := fmt.Sprintf("%08x", crc32.ChecksumIEEE([]byte(field)))
	u := crc[0]
	if !(u >= '0' && u <= '9') {
		return tags[0], strings.ToLower(crc)[:6]
	}
	return tags[0], strings.ToLower(string(u-'0'+'a') + crc[1:6])
}
