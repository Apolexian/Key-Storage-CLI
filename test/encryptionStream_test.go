package test

import (
	"../internal/storage"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestEncryptStream(t *testing.T) {
	logPath, _ := filepath.Abs("../logs/EncryptionStreamTest_Log")
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	v := storage.File("stream-my-key", "stream.secrets")

	err = v.Set("stream-key1", "test-value-1")
	log.Printf("Set stream-key1")
	if err != nil {
		t.Errorf("failed to set: %s", err)
	}

	err = v.Set("stream-key2", "test-value-2")
	log.Printf("Set stream-key2")
	if err != nil {
		t.Errorf("failed to set: %s", err)
	}

	err = v.Set("stream-key3", "test-value-3")
	log.Printf("Set stream-key3")
	if err != nil {
		t.Errorf("failed to set: %s", err)
	}

	plain, err := v.Get("stream-key1")
	if err != nil {
		t.Errorf("failed to get: %s", err)
	}
	log.Printf("Got plain: %s", plain)

	plain, err = v.Get("stream-key2")
	if err != nil {
		t.Errorf("failed to get: %s", err)
	}
	log.Printf("Got plain: %s", plain)

	plain, err = v.Get("stream-key3")
	if err != nil {
		t.Errorf("failed to get: %s", err)
	}
	log.Printf("Got plain: %s", plain)
}
