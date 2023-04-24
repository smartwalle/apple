package internal

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
)

func LoadKeyFromFile(filename string) (*ecdsa.PrivateKey, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return LoadKeyFromBytes(bytes)
}

func LoadKeyFromBytes(bytes []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("must be a valid .p8 PEM file")
	}
	rawKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key, ok := rawKey.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("must be of type ecdsa.PrivateKey")
	}
	return key, nil
}
