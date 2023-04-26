package apple

import (
	"context"
	"encoding/json"
	"github.com/smartwalle/apple/internal"
	"github.com/smartwalle/ngx"
	"io"
	"net/http"
	"strings"
)

const (
	kStoreKitSandboxURL    = "https://api.storekit-sandbox.itunes.apple.com/inApps"
	kStoreKitProductionURL = "https://api.storekit.itunes.apple.com/inApps"
)

type Client struct {
	Client    *http.Client
	token     *internal.Token
	apiDomain string
}

func New(keyfile, keyId, issuer, bundleId string, isProduction bool) (*Client, error) {
	var pKey, err = internal.DecodePrivateKeyFromFile(keyfile)
	if err != nil {
		return nil, err
	}

	var nClient = &Client{}
	nClient.Client = http.DefaultClient
	nClient.token = internal.NewToken(pKey, keyId, issuer, bundleId)

	if isProduction {
		nClient.apiDomain = kStoreKitProductionURL
	} else {
		nClient.apiDomain = kStoreKitSandboxURL
	}

	return nClient, nil
}

func (this *Client) BuildAPI(paths ...string) string {
	var path = this.apiDomain
	for _, p := range paths {
		p = strings.TrimSpace(p)
		if len(p) > 0 {
			if strings.HasSuffix(path, "/") {
				path = path + p
			} else {
				if strings.HasPrefix(p, "/") {
					path = path + p
				} else {
					path = path + "/" + p
				}
			}
		}
	}
	return path
}

func (this *Client) AccessToken() string {
	return this.token.AccessToken()
}

func (this *Client) request(method, url string, param Param, body ngx.Body, result interface{}) (err error) {
	var req = ngx.NewRequest(method, url, ngx.WithClient(this.Client))
	if param != nil {
		req.SetParams(param.Values())
	}
	if body != nil {
		req.SetBody(body)
		req.SetContentType(ngx.ContentTypeJSON)
	}
	req.SetHeader("Authorization", this.AccessToken())

	rsp, err := req.Do(context.Background())
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	switch rsp.StatusCode {
	case http.StatusOK:
		if err = json.Unmarshal(data, &result); err != nil {
			return err
		}
	case http.StatusUnauthorized:
		return &ResponseError{Code: http.StatusUnauthorized, Message: "Unauthenticated"}
	default:
		var rErr *ResponseError
		if err = json.Unmarshal(data, &rErr); err != nil {
			return err
		}
		return rErr
	}
	return nil
}
