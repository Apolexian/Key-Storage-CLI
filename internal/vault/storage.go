package vault

import (
	"../encryption"
	"errors"
)

// Vault is used to store the encoded keys
type Vault struct {
	encodingKey string
	keyValues   map[string]string
}

// Memory will map the encoding string and wrap it into a Vault struct
func Memory(encodingKey string) Vault {
	return Vault{encodingKey: encodingKey,
		keyValues: make(map[string]string)}
}

// Get returns the decrypted version of the key we need
func (v *Vault) Get(key string) (string, error) {
	hex, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("no key value found")
	}
	ret, err := encryption.Decrypt(v.encodingKey, hex)
	if err != nil {
		return "", err
	}

	return ret, nil
}

// Set will load encrypted key into memory
func (v *Vault) Set(key, value string) error {
	encryptedValue, err := encryption.Encrypt(v.encodingKey, value)
	if err != nil {
		return err
	}

	v.keyValues[key] = encryptedValue
	return nil
}
