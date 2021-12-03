package libcsc

import (
	"fmt"

	sha1go "crypto/sha1"
)

func SHA1Sum(data []byte, checksumIn string) (string, bool) {
	checksum := sha1go.Sum(data)
	checksumStr := fmt.Sprintf("%x", checksum)
	return checksumStr, checksumStr == checksumIn
}
