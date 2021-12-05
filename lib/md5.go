package libcsc

import (
	"crypto/md5"
	"fmt"
)

func MD5Sum(data []byte, checksumIn string) (string, bool) {
	checksum := md5.Sum(data)
	checksumStr := fmt.Sprintf("%x", checksum)
	return checksumStr, checksumStr == checksumIn
}
