package test

import (
	"../internal/encryption"
	logger "../internal/logger"
	"math/rand"
	"testing"
	"time"
)

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func RandIntegerMinMax() int {
	rand.Seed(time.Now().UnixNano())
	const MinUint = 0
	// assume that no key/API string is going to be more than 2000 characters
	// can be changed up but causes log file to be larger and may cause overflow
	const MaxInt = 2000
	randomInteger := rand.Intn(MaxInt - MinUint)
	return randomInteger
}
func TestEncrypt(t *testing.T) {

	// To make sure that decryption correctly gets the encrypted string
	// generate random string of random length as key and as tes string
	// encrypt test string, decrypt it and see if correct output
	// run 5 test
	for i := 0; i < 5; i++ {
		testKey := RandStringRunes(RandIntegerMinMax())
		testString := RandStringRunes(RandIntegerMinMax())
		encryptedTest, _ := encryption.Encrypt(testKey, testString)
		decryptedTest, _ := encryption.Decrypt(testKey, encryptedTest)

		logger.TestLogger.Printf("Test Key generated: %s", testKey)
		logger.TestLogger.Printf("String generated: %s", testString)
		logger.TestLogger.Printf("Encrypted Test: %s", encryptedTest)
		logger.TestLogger.Printf("Decrypted Test: %s", decryptedTest)

		if decryptedTest != testString {
			t.Errorf("Incorrect Decryption, expected %s , got %s", testString, decryptedTest)
			logger.TestLogger.Fatalf("Failed test %v", i)
		}
	}
}
