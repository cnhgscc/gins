package minix

import "github.com/cnhgscc/gins/minix/hasher"

var (
	FlagCrc32 = 0
	FlagMD5   = 1
)

func Hasher(flag int, tags ...string) (string, string) {
	switch flag {
	case FlagCrc32:
		return hasher.Crc32Crypto(tags...)
	case FlagMD5:
		return hasher.MD5Crypto(tags...)
	}
	return "", ""
}
