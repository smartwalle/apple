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
	kFetchAuthKeysURL = "https://appleid.apple.com/auth/keys"
)

type IdentityClient struct {
	Client *http.Client
	keys   dbc.Cache[string, *rsa.PublicKey]
	group  singleflight.Group[string]
}

func NewIdentityClient() *IdentityClient {
	var nClient = &IdentityClient{}
	nClient.Client = http.DefaultClient
	nClient.keys = dbc.New[*rsa.PublicKey]()
	nClient.group = singleflight.New()
	return nClient
}

func (this *IdentityClient) DecodeToken(token string) (*User, error) {
	headerBytes, err := base64.RawStdEncoding.DecodeString(strings.Split(token, ".")[0])
	if err != nil {
		return nil, err
	}

	var header *identity.Header
	if err = json.Unmarshal(headerBytes, &header); err != nil {
		return nil, err
	}

	var key = this.GetAuthKey(header.Kid)
	if key == nil {
		return nil, errors.New("invalid token")
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

	this.group.Do(kFetchAuthKeysURL, func(_ string) (interface{}, error) {
		// 从苹果服务器请求 key 数据
		var nKeys, _ = this.requestAuthKeys()

		// 将获取到的 key 缓存起来
		for key, value := range nKeys {
			this.keys.SetEx(key, value, 300) // 300 秒过期
		}
		return nil, nil
	})

	key, _ := this.keys.Get(kid)
	return key
}

// requestAuthKeys https://developer.apple.com/documentation/sign_in_with_apple/fetch_apple_s_public_key_for_verifying_token_signature
func (this *IdentityClient) requestAuthKeys() (map[string]*rsa.PublicKey, error) {
	var req = ngx.NewRequest(http.MethodGet, kFetchAuthKeysURL, ngx.WithClient(this.Client))

	var rsp, err = req.Do(context.Background())
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	var result map[string][]*identity.Key
	if err = json.NewDecoder(rsp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var keys = result["keys"]
	var nKeys = make(map[string]*rsa.PublicKey, len(keys))

	for _, key := range keys {
		var nKey, _ = identity.DecodePublicKey(key)
		if nKey != nil {
			nKeys[key.Kid] = nKey
		}
	}

	return nKeys, nil
}
