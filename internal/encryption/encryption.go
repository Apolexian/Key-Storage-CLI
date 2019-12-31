package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// takes in a key and plaintext and returns the hex representation
// of the the encrypted value
// code is based on standard library examples at:
// https://golang.org/pkg/crypto/cipher/#NewCFBEncrypter
func Encrypt(key, plaintext string) (string, error) {
	block, err := createCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return fmt.Sprintf("%x", ciphertext), nil
}

// takes in a key and the encrypted text(hex) and decrypt it
// code is based on standard library examples at:
// https://golang.org/pkg/crypto/cipher/#NewCFBDecrypter
func Decrypt(key, encryptedText string) (string, error) {
	block, err := createCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("encrypt: cipher is too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil

}

// creates a md5 hash as we need predictable length
// not the most secure, however is okay for purpose
// of simple file storage on local machine
// see : https://golang.org/pkg/crypto/md5/
func createCipher(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}
