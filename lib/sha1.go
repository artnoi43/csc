package libcsc

import (
	"crypto/sha1"
	"fmt"
)

func SHA1Sum(data []byte, checksumIn string) (string, bool) {
	checksum := sha1.Sum(data)
	checksumStr := fmt.Sprintf("%x", checksum)
	return checksumStr, checksumStr == checksumIn
}
