package test

import (
	"../internal/encryption"
	"log"
	"math/rand"
	"os"
	"path/filepath"
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

		// Configuring log file output
		logPath, _ := filepath.Abs("../Logs/EncryptionTest_Log")
		f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		// log test into file in case it is needed at any point
		log.SetOutput(f)
		log.Printf("Testing Encryption test number: %v", i+1)
		log.Printf("Test Key generated: %s", testKey)
		log.Printf("String generated: %s", testString)
		log.Printf("Encrypted Test: %s", encryptedTest)
		log.Printf("Decrypted Test: %s", decryptedTest)
		log.Print("-----------------------------------------")

		if decryptedTest != testString {
			t.Errorf("Incorrect Decryption, expected %s , got %s", testString, decryptedTest)
			log.Fatalf("Failed test %v", i)
		}
	}
}
