package libcsc

type cscFunc func([]byte, string) (string, bool)

var FuncMap = map[Algo]cscFunc{
	SHA256: SHA256Sum,
	SHA1:   SHA1Sum,
	MD5:    MD5Sum,
}
