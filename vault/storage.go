package vault

import "errors"

type Vault struct {
	encodingKey string
	keyValues   map[string]string
}

func Memory(encodingKey string) Vault {
	return Vault{encodingKey: encodingKey,
		keyValues: make(map[string]string)}
}

func (v *Vault) Get(key string) (string, error) {
	if ret, ok := v.keyValues[key]; ok {
		return ret, nil
	}
	return "", errors.New("No key value found.")
}

func (v *Vault) Set(key, value string) error {
	v.keyValues[key] = value
	return nil
}
