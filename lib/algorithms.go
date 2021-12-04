package libcsc

type Algo string
type cscFunc func([]byte, string) (string, bool)

const (
	SHA256 = "SHA256"
	SHA1   = "SHA1"
	MD5    = "MD5"
)

var FuncMap = map[Algo]cscFunc{
	SHA256: SHA256Sum,
	SHA1:   SHA1Sum,
	MD5:    MD5Sum,
}
