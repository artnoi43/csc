package libcsc

import (
	"fmt"

	md5go "crypto/md5"
)

func MD5Sum(data []byte, checksumIn string) (string, bool) {
	checksum := md5go.Sum(data)
	checksumStr := fmt.Sprintf("%x", checksum)
	return checksumStr, checksumStr == checksumIn
}
