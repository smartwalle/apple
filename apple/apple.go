package apple

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	kVerifyReceiptProductionURL = "https://buy.itunes.apple.com/verifyReceipt"
	kVerifyReceiptSandboxURL    = "https://sandbox.itunes.apple.com/verifyReceipt"
)

var (
	ErrBadReceipt = errors.New("bad receipt")
)

type Client struct {
	client *http.Client
}

func New() *Client {
	var c = &Client{}
	c.client = http.DefaultClient
	return c
}

// VerifyReceipt 验证苹果内购交易是否有效，
// 首先请求苹果的服务器，获取票据(receipt)的详细信息，然后验证交易信息(transactionId)是否属于该票据，
// 如果交易信息在票据中，则返回详细的交易信息。
// 注意：本方法会先调用苹果生产环境接口进行票据查询，如果返回票据信息为测试环境中的信息时，则调用测试环境接口进行查询。
func (this *Client) VerifyReceipt(transactionId, receipt string, opts ...Option) (*Trade, *InApp, error) {
	var nOpt = NewOption()
	for _, opt := range opts {
		if opt != nil {
			opt(nOpt)
		}
	}

	var trade, err = this.GetReceipt(receipt, nOpt)
	if err != nil {
		return nil, nil, err
	}

	// 没有交易信息
	if trade == nil {
		return nil, nil, ErrBadReceipt
	}

	// 票据查询失败
	if trade.Status != 0 {
		return nil, nil, fmt.Errorf("bad receipt: %d", trade.Status)
	}

	// 验证 transactionId 和 receipt 是否匹配
	if trade.Receipt != nil {
		for _, info := range trade.Receipt.InApp {
			if info.TransactionId == transactionId {
				return trade, info, nil
			}
		}
	}
	return nil, nil, ErrBadReceipt
}

// GetReceipt 获取票据信息
// 注意：本方法会先调用苹果生产环境接口进行票据查询，如果返回票据信息为测试环境中的信息时，则调用测试环境接口进行查询。
func (this *Client) GetReceipt(receipt string, opts *options) (*Trade, error) {
	// 从生产环境查询
	var trade, err = this.request(kVerifyReceiptProductionURL, receipt, opts)
	if err != nil {
		return nil, err
	}

	// 如果返回票据信息为测试环境中的信息时，则调用测试环境接口进行查询
	if trade != nil && trade.Status == 21007 {
		trade, err = this.request(kVerifyReceiptSandboxURL, receipt, opts)
		if err != nil {
			return nil, err
		}
	}
	return trade, nil
}

func (this *Client) request(url string, receipt string, opts *options) (*Trade, error) {
	var p = &Param{}
	p.Receipt = receipt
	p.Password = opts.password

	var data, err = json.Marshal(p)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	rsp, err := this.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	var trade *Trade
	var decoder = json.NewDecoder(rsp.Body)
	if err = decoder.Decode(&trade); err != nil {
		return nil, err
	}

	return trade, nil
}
