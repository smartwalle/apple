package apple

import (
	"context"
	"encoding/json"
	"github.com/smartwalle/apple/internal"
	"github.com/smartwalle/apple/internal/storekit"
	"github.com/smartwalle/ngx"
	"io"
	"net/http"
	"strings"
)

const (
	kStoreKitSandbox    = "https://api.storekit-sandbox.itunes.apple.com/inApps"
	kStoreKitProduction = "https://api.storekit.itunes.apple.com/inApps"
)

type Client struct {
	Client *http.Client
	token  *storekit.Token
	host   string
}

func New(p8file, keyId, issuer, bundleId string, isProduction bool) (*Client, error) {
	var pKey, err = internal.DecodePrivateKeyFromFile(p8file)
	if err != nil {
		return nil, err
	}

	var nClient = &Client{}
	nClient.Client = http.DefaultClient
	nClient.token = storekit.NewToken(pKey, keyId, issuer, bundleId)

	if isProduction {
		nClient.host = kStoreKitProduction
	} else {
		nClient.host = kStoreKitSandbox
	}

	return nClient, nil
}

func (this *Client) BuildAPI(paths ...string) string {
	var path = this.host
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

func (this *Client) request(method, url string, param Param, body ngx.Body, result interface{}) (err error) {
	var req = ngx.NewRequest(method, url, ngx.WithClient(this.Client))
	if param != nil {
		req.SetParams(param.Values())
	}
	if body != nil {
		req.SetBody(body)
		req.SetContentType(ngx.ContentTypeJSON)
	}
	req.SetHeader("Authorization", this.token.Bearer())

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
