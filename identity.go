package apple

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/smartwalle/apple/internal/identity"
	"github.com/smartwalle/dbc"
	"github.com/smartwalle/ngx"
	"github.com/smartwalle/nsync/singleflight"
	"net/http"
	"strings"
)

const (
	kFetchAuthKeys = "https://appleid.apple.com/auth/keys"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type IdentityClientOptionFunc func(opts *IdentityClient)

// WithKeyExpiration 用于设置从 https://appleid.apple.com/auth/keys 获取的公钥在本地的缓存时间，单位为秒
func WithKeyExpiration(expiration int64) IdentityClientOptionFunc {
	return func(opts *IdentityClient) {
		opts.expiration = expiration
	}
}

type IdentityClient struct {
	Client     *http.Client
	keys       dbc.Cache[string, *rsa.PublicKey]
	group      singleflight.Group[string]
	expiration int64
}

func NewIdentityClient(opts ...IdentityClientOptionFunc) *IdentityClient {
	var nClient = &IdentityClient{}
	nClient.Client = http.DefaultClient
	nClient.keys = dbc.New[*rsa.PublicKey]()
	nClient.group = singleflight.New()
	for _, opt := range opts {
		if opt != nil {
			opt(nClient)
		}
	}
	if nClient.expiration <= 0 {
		nClient.expiration = 300 // 默认 300 秒
	}
	return nClient
}

func (this *IdentityClient) DecodeToken(token string) (*User, error) {
	var payloads = strings.Split(token, ".")
	if len(payloads) < 3 {
		return nil, ErrInvalidToken
	}

	headerBytes, err := base64.RawStdEncoding.DecodeString(payloads[0])
	if err != nil {
		return nil, err
	}

	var header *identity.Header
	if err = json.Unmarshal(headerBytes, &header); err != nil {
		return nil, err
	}

	var key = this.GetAuthKey(header.Kid)
	if key == nil {
		return nil, ErrInvalidToken
	}

	var claims = &identity.Claims{}
	jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) {
		return key, nil
	})

	var nUser = &User{}
	nUser.Id = claims.Subject
	nUser.BundleId = strings.Join(claims.Audience, ";")
	nUser.AuthTime = claims.AuthTime
	nUser.Email = claims.Email
	nUser.EmailVerified = claims.EmailVerified
	nUser.IsPrivateEmail = claims.IsPrivateEmail
	nUser.RealUserStatus = claims.RealUserStatus
	return nUser, nil
}

func (this *IdentityClient) GetAuthKey(kid string) *rsa.PublicKey {
	// 从本地缓存中查询 key 信息，存在则直接返回
	if key, _ := this.keys.Get(kid); key != nil {
		return key
	}

	this.group.Do(kFetchAuthKeys, func(_ string) (interface{}, error) {
		// 从苹果服务器请求 key 数据
		var nKeys, _ = this.requestAuthKeys()

		for _, key := range nKeys {
			var nKey, _ = identity.DecodePublicKey(key.N, key.E)
			if nKey != nil {
				this.keys.SetEx(key.Kid, nKey, this.expiration)
			}
		}
		return nil, nil
	})

	key, _ := this.keys.Get(kid)
	return key
}

// requestAuthKeys https://developer.apple.com/documentation/sign_in_with_apple/fetch_apple_s_public_key_for_verifying_token_signature
func (this *IdentityClient) requestAuthKeys() ([]identity.Key, error) {
	var req = ngx.NewRequest(http.MethodGet, kFetchAuthKeys, ngx.WithClient(this.Client))

	var rsp, err = req.Do(context.Background())
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	var aux = &struct {
		Keys []identity.Key `json:"keys"`
	}{}
	if err = json.NewDecoder(rsp.Body).Decode(&aux); err != nil {
		return nil, err
	}

	return aux.Keys, nil
}
