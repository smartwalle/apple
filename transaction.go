package apple

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	kTransaction            = "/v1/transactions/"
	kTransactionHistory     = "/v1/history/"
	kTransactionConsumption = "/v1/transactions/consumption/"
)

// GetTransaction https://developer.apple.com/documentation/appstoreserverapi/get_transaction_info
func (c *Client) GetTransaction(transactionId string) (result *TransactionResponse, err error) {
	err = c.request(http.MethodGet, c.BuildAPI(kTransaction, transactionId), nil, nil, &result)
	return result, err
}

// GetTransactionHistory https://developer.apple.com/documentation/appstoreserverapi/get_transaction_history
func (c *Client) GetTransactionHistory(transactionId string, param TransactionHistoryParam) (result *TransactionHistoryResponse, err error) {
	err = c.request(http.MethodGet, c.BuildAPI(kTransactionHistory, transactionId), param, nil, &result)
	return result, err
}

// SendConsumptionInformation https://developer.apple.com/documentation/appstoreserverapi/send_consumption_information
func (c *Client) SendConsumptionInformation(transactionId string, param ConsumptionParam) (err error) {
	data, err := json.Marshal(param)
	if err != nil {
		return err
	}
	err = c.request(http.MethodPut, c.BuildAPI(kTransactionConsumption, transactionId), nil, bytes.NewReader(data), nil)
	return err
}
