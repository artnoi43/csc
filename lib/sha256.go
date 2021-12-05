package libcsc

import (
	"crypto/sha256"
	"fmt"
)

func SHA256Sum(data []byte, checksumIn string) (string, bool) {
	checksum := sha256.Sum256(data)
	checksumStr := fmt.Sprintf("%x", checksum)
	return checksumStr, checksumStr == checksumIn
}
