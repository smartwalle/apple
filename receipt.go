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

// VerifyReceipt 验证苹果内购交易是否有效 https://developer.apple.com/documentation/appstorereceipts/verifyreceipt
// 首先请求苹果的服务器，获取票据(receipt)的详细信息，然后验证交易信息(transactionId)是否属于该票据，
// 如果交易信息在票据中，则返回详细的交易信息。
// 注意：本方法会先调用苹果生产环境接口进行票据查询，如果返回票据信息为测试环境中的信息时，则调用测试环境接口进行查询。
func VerifyReceipt(transactionId, receipt string, opts ...VerifyReceiptOptionFunc) (*ReceiptResponse, *InApp, error) {
	var summary, err = GetReceipt(receipt, opts...)
	if err != nil {
		return nil, nil, err
	}

	// 没有交易信息
	if summary == nil {
		return nil, nil, ErrBadReceipt
	}

	// 票据查询失败
	if summary.Status != 0 {
		return nil, nil, fmt.Errorf("bad receipt: %d", summary.Status)
	}

	// 验证 transactionId 和 receipt 是否匹配
	if summary.Receipt != nil {
		for _, info := range summary.Receipt.InApp {
			if info.TransactionId == transactionId {
				return summary, info, nil
			}
		}
	}
	return nil, nil, ErrBadReceipt
}

// GetReceipt 获取票据信息
// 注意：本方法会先调用苹果生产环境接口进行票据查询，如果返回票据信息为测试环境中的信息时，则调用测试环境接口进行查询。
func GetReceipt(receipt string, opts ...VerifyReceiptOptionFunc) (*ReceiptResponse, error) {
	var nOpt = &ReceiptOptions{}
	nOpt.Receipt = receipt
	for _, opt := range opts {
		if opt != nil {
			opt(nOpt)
		}
	}

	// 从生产环境查询
	var summary, err = getReceipt(kVerifyReceiptProductionURL, nOpt)
	if err != nil {
		return nil, err
	}

	// 如果返回票据信息为测试环境中的信息时，则调用测试环境接口进行查询
	if summary != nil && summary.Status == 21007 {
		summary, err = getReceipt(kVerifyReceiptSandboxURL, nOpt)
		if err != nil {
			return nil, err
		}
	}
	return summary, nil
}

func getReceipt(url string, opts *ReceiptOptions) (*ReceiptResponse, error) {
	var data, err = json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	var summary *ReceiptResponse
	var decoder = json.NewDecoder(rsp.Body)
	if err = decoder.Decode(&summary); err != nil {
		return nil, err
	}

	return summary, nil
}
