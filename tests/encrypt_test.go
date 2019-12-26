package tests

import (
	"../encryption"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestEncrypt(t *testing.T) {

	// To make sure that decryption correctly gets the encrypted string
	for i := 1; i < 100; i++ {
		testKey := RandStringRunes(i * 10)
		testString := RandStringRunes(i)
		encryptedTest, _ := encryption.Encrypt(testKey, testString)
		decryptedTest, _ := encryption.Decrypt(testKey, encryptedTest)

		// Configuring log file output
		logPath,_ := filepath.Abs("../Logs/EncryptionTest_Log")
		f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()

		// log test into file in case it is needed at any point
		log.SetOutput(f)
		log.Printf("Testing Encryption test number: %v", i)
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
