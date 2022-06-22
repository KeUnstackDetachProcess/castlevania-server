package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func NewKey() ([]byte, error) {

	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	fmt.Println(key)
	return key, nil
}

func NewIv() ([]byte, error) {

	iv := make([]byte, 16)

	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	fmt.Println(iv)
	return iv, nil
}

func Encrypt(data []byte, key []byte, iv []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	pkcs5 := PKCS5Padding(data, block.BlockSize())
	ciphertext := make([]byte, len(pkcs5))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, pkcs5)

	hexData := make([]byte, hex.EncodedLen(len(ciphertext)))
	hex.Encode(hexData, ciphertext)

	return hexData, nil
}

func Decrypt(encoded_ciphertext []byte, key []byte, iv []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decodedData := make([]byte, hex.DecodedLen(len(encoded_ciphertext)))
	hex.Decode(decodedData, encoded_ciphertext)

	data := make([]byte, len(decodedData))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, decodedData)

	return PKCS5Unpadding(data), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
