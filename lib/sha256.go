package libcsc

import (
	"fmt"

	sha256go "crypto/sha256"
)

func SHA256Sum(data []byte, checksumIn string) (string, bool) {
	checksum := sha256go.Sum256(data)
	checksumStr := fmt.Sprintf("%x", checksum)
	return checksumStr, checksumStr == checksumIn
}
