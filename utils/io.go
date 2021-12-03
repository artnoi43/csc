package utils

import (
	"bytes"
	"io"
	"os"

	"github.com/pkg/errors"
)

func ReadFile(path string) ([]byte, error) {
	var buf = new(bytes.Buffer)
	fp, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}
	defer fp.Close()
	_, err = io.Copy(buf, fp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}
	return buf.Bytes(), nil
}
