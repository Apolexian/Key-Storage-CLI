package test

import (
	logger "../internal/logger"
	"../internal/storage"
	"testing"
)

func TestStorage(t *testing.T) {
	v := storage.File("test-key", "storage.secrets")

	err := v.Set("key1", "test-value-1")
	logger.TestLogger.Printf("Set key1")
	if err != nil {
		t.Errorf("failed to set: %s", err)
		logger.TestLogger.Printf("error on setting key1: %s", err)
	}

	err = v.Set("key2", "test-value-2")
	logger.TestLogger.Printf("Set key2")
	if err != nil {
		t.Errorf("failed to set: %s", err)
		logger.TestLogger.Printf("error on setting key2: %s", err)
	}

	err = v.Set("key3", "test-value-3")
	logger.TestLogger.Printf("Set key3")
	if err != nil {
		t.Errorf("failed to set: %s", err)
		logger.TestLogger.Printf("error on setting key3: %s", err)
	}

	plain, err := v.Get("key1")
	logger.TestLogger.Printf("Got plain: %s", plain)
	if err != nil {
		t.Errorf("failed to get demo key with error: %s", err)
		logger.TestLogger.Printf("error on getting key1: %s", err)
	}
	if plain != "test-value-1" {
		t.Errorf("expected test-value-1, got: %s", plain)
		logger.TestLogger.Printf("expected test-value-1, got: %s", plain)
	}

	plain, err = v.Get("key2")
	logger.TestLogger.Printf("Got plain: %s", plain)
	if err != nil {
		t.Errorf("failed to get demo key with error: %s", err)
		logger.TestLogger.Printf("error on getting key2: %s", err)
	}
	if plain != "test-value-2" {
		t.Errorf("expected test-value-2, got: %s", plain)
		logger.TestLogger.Printf("expected test-value-2, got: %s", plain)
	}

	plain, err = v.Get("key3")
	logger.TestLogger.Printf("Got plain: %s", plain)
	if err != nil {
		t.Errorf("failed to get demo key with error: %s", err)
		logger.TestLogger.Printf("error on getting key3: %s", err)
	}
	if plain != "test-value-3" {
		t.Errorf("expected test-value-3, got: %s", plain)
		logger.TestLogger.Printf("expected test-value-3, got: %s", plain)
	}
}
