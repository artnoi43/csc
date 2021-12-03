package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	csc "github.com/artnoi43/csc/lib"
	"github.com/artnoi43/csc/utils"
)

type flags struct {
	algo     string
	filename string
	checksum string
}

func (f *flags) parse() {
	flag.StringVar(&f.algo, "a", "SHA256", "hash algorithm")
	flag.StringVar(&f.filename, "f", "", "file name")
	flag.StringVar(&f.checksum, "c", "", "checksum string to check against")
	flag.Parse()
}

var (
	f    flags
	algo csc.Algo
)

func init() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(0)
	}

	f.parse()
	if f.filename == "" || f.checksum == "" {
		usage()
	}
	if strings.EqualFold(f.algo, "SHA256") {
		algo = csc.SHA256
	} else if strings.EqualFold(f.algo, "SHA1") {
		algo = csc.SHA1
	} else if strings.EqualFold(f.algo, "MD5") {
		algo = csc.MD5
	}
}

func main() {
	data, err := utils.ReadFile(f.filename)
	if err != nil {
		os.Stderr.Write([]byte(fmt.Sprintf("Error reading read file: %s\n", err.Error())))
		os.Exit(2)
	}

	checksumFunc := csc.FuncMap[algo]
	fileChecksum, matched := checksumFunc(data, f.checksum)

	os.Stdout.Write([]byte(fmt.Sprintf("[%v] In:\t%s\n", algo, f.checksum)))
	os.Stdout.Write([]byte(fmt.Sprintf("[%v] File:\t%s\n", algo, fileChecksum)))
	if matched {
		os.Stdout.Write([]byte(fmt.Sprintf("[%v] Checksum matched\n", algo)))
		os.Exit(0)
	} else {
		os.Stdout.Write([]byte(fmt.Sprintf("[%v] Checksum not matched\n", algo)))
		os.Exit(1)
	}
}

func usage() {
	os.Stderr.Write([]byte(usageStr))
}

const (
	usageStr = `
csc - a checksum checker
usage:
csc -a <algorithm> -f <filename> -c <checksum to compare>

`
)
