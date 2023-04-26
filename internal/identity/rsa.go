package identity

import (
	"crypto/rsa"
	"encoding/base64"
	"math/big"
)

func DecodePublicKey(key *Key) (*rsa.PublicKey, error) {
	nByte, err := base64.RawURLEncoding.DecodeString(key.N)
	if err != nil {
		return nil, err
	}
	eByte, err := base64.RawURLEncoding.DecodeString(key.E)
	if err != nil {
		return nil, err
	}

	var pKey rsa.PublicKey

	pKey.N = big.NewInt(0).SetBytes(nByte)
	pKey.E = int(big.NewInt(0).SetBytes(eByte).Uint64())

	return &pKey, nil
}
