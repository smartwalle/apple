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

	accessToken string

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

func (this *Token) generate() (string, int64, error) {
	var issuedAt = time.Now().Unix()
	var expiredAt = issuedAt + kTokenTimeout
	var nToken = jwt.New(jwt.SigningMethodES256)
	nToken.Header["kid"] = this.keyId
	nToken.Claims = jwt.MapClaims{
		"iss":   this.issuer,
		"iat":   issuedAt,
		"exp":   expiredAt,
		"aud":   "appstoreconnect-v1",
		"nonce": uuid.NewString(),
		"bid":   this.bundleId,
	}
	var bearer, err = nToken.SignedString(this.privateKey)
	if err != nil {
		return "", 0, err
	}
	return fmt.Sprintf("Bearer %s", bearer), issuedAt, nil
}

func (this *Token) AccessToken() string {
	this.Lock()
	defer this.Unlock()

	if this.expired() {
		this.accessToken, this.issuedAt, _ = this.generate()
	}
	return this.accessToken
}

func (this *Token) expired() bool {
	return time.Now().Unix() >= (this.issuedAt + kTokenTimeout)
}
