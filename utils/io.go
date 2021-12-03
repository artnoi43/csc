package utils

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

func ReadFile(path string) ([]byte, error) {
	var buf = new(bytes.Buffer)
	fp, err := os.Open(path)
	if err != nil {
		var emptyPath string
		if len(path) == 0 {
			emptyPath = "\nReminder: specify path with -f\n"
		}
		return nil, errors.Wrap(err, fmt.Sprintf("%sfailed to open file '%s'", emptyPath, path))
	}
	defer fp.Close()
	_, err = io.Copy(buf, fp)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to read file '%s'", path))
	}
	return buf.Bytes(), nil
}
