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
func (this *Client) GetTransaction(transactionId string) (result *TransactionResponse, err error) {
	err = this.request(http.MethodGet, this.BuildAPI(kTransaction, transactionId), nil, nil, &result)
	return result, err
}

// GetTransactionHistory https://developer.apple.com/documentation/appstoreserverapi/get_transaction_history
func (this *Client) GetTransactionHistory(transactionId string, param TransactionHistoryParam) (result *TransactionHistoryResponse, err error) {
	err = this.request(http.MethodGet, this.BuildAPI(kTransactionHistory, transactionId), param, nil, &result)
	return result, err
}

// SendConsumptionInformation https://developer.apple.com/documentation/appstoreserverapi/send_consumption_information
func (this *Client) SendConsumptionInformation(transactionId string, param ConsumptionParam) (err error) {
	data, err := json.Marshal(param)
	if err != nil {
		return err
	}
	err = this.request(http.MethodPut, this.BuildAPI(kTransactionConsumption, transactionId), nil, bytes.NewReader(data), nil)
	return err
}
