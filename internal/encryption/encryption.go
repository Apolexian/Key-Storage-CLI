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

// encryptStream provides a stream wrapper that will encrypt the provider
func encryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := createCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewCFBEncrypter(block, iv), nil
}

// Encrypt takes in a key and plaintext and returns the hex representation
// of the the encrypted value
// code is based on standard library examples at:
// https://golang.org/pkg/crypto/cipher/#NewCFBEncrypter
func Encrypt(key, plaintext string) (string, error) {
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream, err := encryptStream(key, iv)
	if err != nil {
		return "", err
	}
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))
	return fmt.Sprintf("%x", ciphertext), nil
}

// EncryptWriter passes a writer to make sure that whatever
// is written to the file is encrypted
func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream, err := encryptStream(key, iv)
	if err != nil {
		return nil, err
	}
	n, err := w.Write(iv)
	if n != len(iv) || err != nil {
		return nil, errors.New("encrypt writer not able to write full iv to writer")
	}
	return &cipher.StreamWriter{S: stream, W: w}, nil
}

// decryptStream will decrypt the provided encrypted stream wrapper and return a
// decrypted stream
func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := createCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewCFBDecrypter(block, iv), nil

}

// Decrypt takes in a key and the encrypted text(hex) and decrypt it
// code is based on standard library examples at:
// https://golang.org/pkg/crypto/cipher/#NewCFBDecrypter
func Decrypt(key, encryptedText string) (string, error) {
	ciphertext, err := hex.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("decrypt: cipher is too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream, err := decryptStream(key, iv)
	if err != nil {
		return "", err
	}
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}

// DecryptReader will return a wrapper reader that will decrypt data from original reader
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if n < len(iv) || err != nil {
		errors.New("decrypt unable to read full iv")
	}
	stream, err := decryptStream(key, iv)
	if err != nil {
		return nil, err
	}
	return &cipher.StreamReader{S: stream, R: r}, nil
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
