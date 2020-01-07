package vault

import (
	"../encryption"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// Vault is used to store the encoded keys
type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

// File will map the encoding string and wrap it into a Vault struct
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
	defer f.Close()
	var sb strings.Builder
	_, err = io.Copy(&sb, f)
	if err != nil {
		return err
	}
	decryptedJSON, err := encryption.Decrypt(v.encodingKey, sb.String())
	if err != nil {
		return err
	}
	r := strings.NewReader(decryptedJSON)
	dec := json.NewDecoder(r)
	err = dec.Decode(&v.keyValues)
	if err != nil {
		return err
	}
	return nil
}

// putKeyValues will put the encrypted values into a JSON
// and save to file
func (v *Vault) putKeyValues() error {
	var sb strings.Builder
	enc := json.NewEncoder(&sb)
	err := enc.Encode(v.keyValues)
	if err != nil {
		return err
	}
	encryptedJSON, err := encryption.Encrypt(v.encodingKey, sb.String())
	if err != nil {
		return err
	}
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprint(f, encryptedJSON)
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
