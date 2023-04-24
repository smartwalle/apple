package apple

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	kTransactionHistory     = "/v1/history/"
	kTransactionConsumption = "/v1/transactions/consumption/"
)

// GetTransactionHistory https://developer.apple.com/documentation/appstoreserverapi/get_transaction_history
func (this *Client) GetTransactionHistory(transactionId string, param TransactionHistoryParam) (result *TransactionHistoryRsp, err error) {
	err = this.request(http.MethodGet, this.BuildAPI(kTransactionHistory, transactionId), param, nil, &result)
	return result, err
}

// SendConsumptionInformation https://developer.apple.com/documentation/appstoreserverapi/send_consumption_information
func (this *Client) SendConsumptionInformation(transactionId string, param ConsumptionParam) (result *ConsumptionRsp, err error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	err = this.request(http.MethodPut, this.BuildAPI(kTransactionConsumption, transactionId), nil, bytes.NewReader(data), &result)
	return result, err
}
