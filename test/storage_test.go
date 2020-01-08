package test

import (
	"../internal/vault"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestStorage(t *testing.T) {

	logPath, _ := filepath.Abs("../logs/StorageTest_Log")
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	v := vault.File("test-key", ".secrets")

	err = v.Set("key1", "test-value-1")
	log.Printf("Set key1")
	if err != nil {
		t.Errorf("failed to set: %s", err)
		log.Printf("error on setting key1: %s", err)
	}

	err = v.Set("key2", "test-value2")
	log.Printf("Set key2")
	if err != nil {
		t.Errorf("failed to set: %s", err)
		log.Printf("error on setting key2: %s", err)
	}

	err = v.Set("key3", "test-value3")
	log.Printf("Set key3")
	if err != nil {
		t.Errorf("failed to set: %s", err)
		log.Printf("error on setting key3: %s", err)
	}

	plain, err := v.Get("key1")
	log.Printf("Got plain: %s", plain)
	if err != nil {
		t.Errorf("failed to get demo key with error: %s", err)
		log.Printf("error on getting key1: %s", err)
	}

	plain, err = v.Get("key2")
	log.Printf("Got plain: %s", plain)
	if err != nil {
		t.Errorf("failed to get demo key with error: %s", err)
		log.Printf("error on getting key2: %s", err)
	}

	plain, err = v.Get("key3")
	log.Printf("Got plain: %s", plain)
	if err != nil {
		t.Errorf("failed to get demo key with error: %s", err)
		log.Printf("error on getting key3: %s", err)
	}
}
