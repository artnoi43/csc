package libcsc

type cscFunc func([]byte, string) (string, bool)

var FuncMap = map[Algo]cscFunc{
	sha256: SHA256Sum,
	sha1:   SHA1Sum,
	md5:    MD5Sum,
}
