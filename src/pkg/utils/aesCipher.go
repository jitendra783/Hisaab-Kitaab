package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

type aesCipher struct {
	SecretKey   []byte
	BlockSize   int
	IsIvPresent bool
}

type AesCipherGroup interface {
	AuthTokenDecryption(encryptedData string) (string, error)
}

func NewAesCipherService(secretKey string, isIvPresent bool) AesCipherGroup {
	return &aesCipher{
		SecretKey:   []byte(secretKey),
		BlockSize:   aes.BlockSize,
		IsIvPresent: isIvPresent,
	}
}

func (a *aesCipher) unPad(data []byte) []byte {
	unpadding := int(data[len(data)-1])
	return data[:(len(data) - unpadding)]
}

func (a *aesCipher) AuthTokenDecryption(encryptedData string) (string, error) {
	iv := make([]byte, 16)
	encryptedByteData := []byte(encryptedData)
	//Take out IV from encrypted access toke
	if a.IsIvPresent {
		ivPosition := len(encryptedData) - 16
		iv = encryptedByteData[ivPosition:]
		encryptedByteData = encryptedByteData[:ivPosition]
	}

	data, err := base64.StdEncoding.DecodeString(string(encryptedByteData))
	if err != nil {
		return "", fmt.Errorf("error in decoding the string, error: %v", err)
	}

	if len(data) < a.BlockSize {
		return "", fmt.Errorf("error in decoding the string, encrypted too short")
	}

	// CBC mode always works in whole blocks.
	if len(data)%a.BlockSize != 0 {
		return "", fmt.Errorf("error in decoding the string, ciphertext is not a multiple of the block size")
	}

	cipherBlock, err := aes.NewCipher(a.SecretKey)
	if err != nil {
		return "", fmt.Errorf("error in creating new cipher, Error: %v", err)
	}

	decryptor := cipher.NewCBCDecrypter(cipherBlock, iv)

	decryptedBytes := make([]byte, len(data))
	decryptor.CryptBlocks(decryptedBytes, []byte(data))

	decryptedBytes = a.unPad(decryptedBytes)

	return string(decryptedBytes), nil
}
