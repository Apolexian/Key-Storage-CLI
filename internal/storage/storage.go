package storage

import (
	"../encryption"
	"encoding/json"
	"errors"
	"io"
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
	r, err := encryption.DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.readKeyValues(r)
}

// readKeyValues will read and decode json values from file
func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

// putKeyValues will put the encrypted values into a JSON
// and save to file
func (v *Vault) putKeyValues() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	w, err := encryption.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}
	return v.writeKeyValues(w)
}

// writeKeyValues encodes io writer wrapper to json
func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v.keyValues)
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
	err := v.getKeyValues()
	if err != nil {
		return err
	}
	v.keyValues[key] = value
	err = v.putKeyValues()
	return err
}
