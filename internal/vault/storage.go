package vault

import (
	"../encryption"
	"encoding/json"
	"errors"
	"os"
	"sync"
)

// Vault is used to store the encoded keys
type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

// Memory will map the encoding string and wrap it into a Vault struct
func File(encodingKey, filepath string) *Vault {
	return &Vault{encodingKey: encodingKey,
		keyValues: make(map[string]string),
		filepath:  filepath}
}

// getKeyValues returns the values from the json file
// if not present makes new one instead of error
func (v *Vault) getKeyValues() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&v.keyValues)
	if err != nil {
		return err
	}
	return nil
}

// Get returns the decrypted version of the key we need
func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.getKeyValues()
	if err != nil {
		return "", err
	}

	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("no key value found")
	}
	return value, nil
}

// Set will load encrypted key into memory
func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	encryptedValue, err := encryption.Encrypt(v.encodingKey, value)
	if err != nil {
		return err
	}
	err = v.getKeyValues()
	if err != nil {
		return err
	}
	v.keyValues[key] = encryptedValue
	err = v.putKeyValues()
	return err
}
