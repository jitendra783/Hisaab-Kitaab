package middleware
type aesCipher struct {
	SecretKey   []byte
	BlockSize   int
	IsIvPresent bool
}