package storekit

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"sync"
	"time"
)

const (
	kTokenTimeout = 3000
)

type Token struct {
	sync.Mutex

	bearer     string
	issuedAt   int64
	keyId      string
	issuer     string
	bundleId   string
	privateKey *ecdsa.PrivateKey
}

func NewToken(privateKey *ecdsa.PrivateKey, keyId, issuer, bundleId string) *Token {
	return &Token{
		keyId:      keyId,
		issuer:     issuer,
		bundleId:   bundleId,
		privateKey: privateKey,
	}
}

func (t *Token) generate() (string, int64, error) {
	var issuedAt = time.Now().Unix()
	var expiredAt = issuedAt + kTokenTimeout
	var nToken = jwt.New(jwt.SigningMethodES256)
	nToken.Header["kid"] = t.keyId
	nToken.Claims = jwt.MapClaims{
		"iss":   t.issuer,
		"iat":   issuedAt,
		"exp":   expiredAt,
		"aud":   "appstoreconnect-v1",
		"nonce": uuid.NewString(),
		"bid":   t.bundleId,
	}
	var bearer, err = nToken.SignedString(t.privateKey)
	if err != nil {
		return "", 0, err
	}
	return fmt.Sprintf("Bearer %s", bearer), issuedAt, nil
}

func (t *Token) Bearer() string {
	t.Lock()
	defer t.Unlock()

	if t.expired() {
		t.bearer, t.issuedAt, _ = t.generate()
	}
	return t.bearer
}

func (t *Token) expired() bool {
	return time.Now().Unix() >= (t.issuedAt + kTokenTimeout)
}
