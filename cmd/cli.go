package main

import "flag"

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
