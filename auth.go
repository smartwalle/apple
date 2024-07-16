package apple

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/smartwalle/apple/internal/auth"
	"github.com/smartwalle/dbc"
	"github.com/smartwalle/ngx"
	"github.com/smartwalle/nsync/singleflight"
	"net/http"
	"strings"
	"time"
)

const (
	kFetchAuthKeys = "https://appleid.apple.com/auth/keys"
	kIssuer        = "https://appleid.apple.com"
)

var (
	ErrInvalidToken    = errors.New("invalid token")
	ErrInvalidIssuer   = errors.New("invalid issuer")
	ErrInvalidBundleId = errors.New("invalid bundle id")
	ErrTokenExpired    = errors.New("token is expired")
)

type AuthClientOption func(c *AuthClient)

// WithKeyExpiration 用于设置从 https://appleid.apple.com/auth/keys 获取的公钥在本地的缓存时间，单位为秒
func WithKeyExpiration(expiration int64) AuthClientOption {
	return func(c *AuthClient) {
		c.expiration = expiration
	}
}

// WithBundleId 用于设置 VerifyToken() 方法需要的 BundleId 信息
func WithBundleId(bundleId string) AuthClientOption {
	return func(c *AuthClient) {
		c.bundleId = bundleId
	}
}

// AuthClient 苹果登录验证
type AuthClient struct {
	Client     *http.Client
	keys       dbc.Cache[string, *rsa.PublicKey]
	group      singleflight.Group[string, interface{}]
	expiration int64
	bundleId   string
}

func NewAuthClient(opts ...AuthClientOption) *AuthClient {
	var nClient = &AuthClient{}
	nClient.Client = http.DefaultClient
	nClient.keys = dbc.New[*rsa.PublicKey]()
	nClient.group = singleflight.New[interface{}]()
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

// DecodeToken 解析 Token
//
// 只对 Token 进行解析，不会验证合法性
func (c *AuthClient) DecodeToken(token string) (*User, error) {
	var payloads = strings.Split(token, ".")
	if len(payloads) < 3 {
		return nil, ErrInvalidToken
	}

	headerBytes, err := base64.RawStdEncoding.DecodeString(payloads[0])
	if err != nil {
		return nil, err
	}

	var header *auth.Header
	if err = json.Unmarshal(headerBytes, &header); err != nil {
		return nil, err
	}

	var key = c.GetAuthKey(header.Kid)
	if key == nil {
		return nil, ErrInvalidToken
	}

	var claims = &auth.Claims{}
	jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) {
		return key, nil
	})

	var user = &User{}
	user.Id = claims.Subject
	user.Issuer = claims.Issuer
	user.BundleId = strings.Join(claims.Audience, ";")
	user.Email = claims.Email
	user.EmailVerified = bool(claims.EmailVerified)
	user.IsPrivateEmail = bool(claims.IsPrivateEmail)
	user.RealUserStatus = claims.RealUserStatus
	user.TransferSub = claims.TransferSub
	user.Nonce = claims.Nonce
	user.AuthTime = int64(claims.AuthTime)
	user.IssuedAt = claims.IssuedAt.Unix()
	user.ExpiresAt = claims.ExpiresAt.Unix()
	return user, nil
}

// VerifyToken 解析并验证 Token https://developer.apple.com/documentation/sign_in_with_apple/sign_in_with_apple_rest_api/verifying_a_user#3383769
//
// 会对 Token 的合法性进行验证，主要判断 BundleId 和 Issuer 是否正确以及 Token 是否在有效期内
func (c *AuthClient) VerifyToken(token string) (*User, error) {
	var user, err = c.DecodeToken(token)
	if err != nil {
		return nil, err
	}

	if user.BundleId != c.bundleId {
		return nil, ErrInvalidBundleId
	}

	if user.Issuer != kIssuer {
		return nil, ErrInvalidIssuer
	}

	if user.ExpiresAt <= time.Now().Unix() {
		return nil, ErrTokenExpired
	}

	return user, nil
}

func (c *AuthClient) GetAuthKey(kid string) *rsa.PublicKey {
	// 从本地缓存中查询 key 信息，存在则直接返回
	if key, _ := c.keys.Get(kid); key != nil {
		return key
	}

	c.group.Do(kFetchAuthKeys, func(_ string) (interface{}, error) {
		// 从苹果服务器请求 key 数据
		var nKeys, _ = c.requestAuthKeys()

		for _, key := range nKeys {
			var nKey, _ = auth.DecodePublicKey(key.N, key.E)
			if nKey != nil {
				c.keys.SetEx(key.Kid, nKey, c.expiration)
			}
		}
		return nil, nil
	})

	key, _ := c.keys.Get(kid)
	return key
}

// requestAuthKeys https://developer.apple.com/documentation/sign_in_with_apple/fetch_apple_s_public_key_for_verifying_token_signature
func (c *AuthClient) requestAuthKeys() ([]auth.Key, error) {
	var req = ngx.NewRequest(http.MethodGet, kFetchAuthKeys, ngx.WithClient(c.Client))

	var rsp, err = req.Do(context.Background())
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	var aux = struct {
		Keys []auth.Key `json:"keys"`
	}{}
	if err = json.NewDecoder(rsp.Body).Decode(&aux); err != nil {
		return nil, err
	}

	return aux.Keys, nil
}
