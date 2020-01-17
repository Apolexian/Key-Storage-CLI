package test

import (
	"../internal/logger"
	"../internal/storage"
	"testing"
)

func TestEncryptStream(t *testing.T) {
	v := storage.File("stream-my-key", "stream.secrets")

	err := v.Set("stream-key1", "test-value-1")
	logger.TestLogger.Printf("Set stream-key1")
	if err != nil {
		t.Errorf("failed to set: %s", err)
		logger.TestLogger.Printf("failed to set: %s", err)
	}

	err = v.Set("stream-key2", "test-value-2")
	logger.TestLogger.Printf("Set stream-key2")
	if err != nil {
		t.Errorf("failed to set: %s", err)
		logger.TestLogger.Printf("failed to set: %s", err)
	}

	err = v.Set("stream-key3", "test-value-3")
	logger.TestLogger.Printf("Set stream-key3")
	if err != nil {
		t.Errorf("failed to set: %s", err)
		logger.TestLogger.Printf("failed to set: %s", err)
	}

	plain, err := v.Get("stream-key1")
	if err != nil {
		t.Errorf("failed to get: %s", err)
		logger.TestLogger.Printf("failed to get: %s", err)
	}
	logger.TestLogger.Printf("Got plain: %s", plain)

	plain, err = v.Get("stream-key2")
	if err != nil {
		t.Errorf("failed to get: %s", err)
		logger.TestLogger.Printf("failed to get: %s", err)
	}
	logger.TestLogger.Printf("Got plain: %s", plain)

	plain, err = v.Get("stream-key3")
	if err != nil {
		t.Errorf("failed to get: %s", err)
		logger.TestLogger.Printf("failed to get: %s", err)
	}
	logger.TestLogger.Printf("Got plain: %s", plain)
}
