package tripleDES

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func tripleDES() {
	K1, K2, K3 := "19961997", "19981999", "20002001"

	triple_keys := K1 + K2 + K3
	plain_text := []byte("Golang is Awesome")

	encrypted, err := TripleDESEncryption(plain_text, []byte(triple_keys))

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s encrypt to %x \n", string(plain_text[:]), string(encrypted[:]))

	decrypted, err := TripleDESDecryption(encrypted, []byte(triple_keys))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s decrypted to %x \n", encrypted, decrypted)

}

func TripleDESEncryption(data, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	cipher_text := key
	iv := cipher_text[:des.BlockSize]
	origData := PKCS5Padding(data, block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(origData))
	mode.CryptBlocks(encrypted, origData)
	return encrypted, nil
}

func TripleDESDecryption(data, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := key
	iv := ciphertext[:des.BlockSize]

	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(data))
	decrypter.CryptBlocks(decrypted, data)
	decrypted = PKCS5UnPadding(decrypted)
	return decrypted, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
