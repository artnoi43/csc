package main

import (
	"flag"
	"fmt"
	"os"

	csc "github.com/artnoi43/csc/lib"
	"github.com/artnoi43/csc/utils"
)

type flags struct {
	algo     string
	filename string
	checksum string
}

func (f *flags) parse() {
	flag.StringVar(&f.algo, "a", "sha256", "Algorithm")
	flag.StringVar(&f.filename, "f", "", "File name")
	flag.StringVar(&f.checksum, "c", "", "Checksum string to check against")
	flag.Parse()
}

var f flags

func init() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(0)
	}
	f.parse()
	if f.filename == "" || f.checksum == "" {
		usage()
	}
}

func main() {
	data, err := utils.ReadFile(f.filename)
	if err != nil {
		os.Stderr.Write([]byte(fmt.Sprintf("Error reading read file: %s\n", err.Error())))
		os.Exit(2)
	}

	checksumFunc := csc.FuncMap[csc.Algo(f.algo)]
	fileChecksum, matched := checksumFunc(data, f.checksum)

	os.Stdout.Write([]byte(fmt.Sprintf("In:\t%s\n", f.checksum)))
	os.Stdout.Write([]byte(fmt.Sprintf("File:\t%s\n", fileChecksum)))
	if matched {
		os.Stdout.Write([]byte("Checksum matched\n"))
		os.Exit(0)
	} else {
		os.Stdout.Write([]byte("Checksum not matched\n"))
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
