package apple

import (
	"net/http"
)

const (
	kTransactionHistory = "/v1/history/"
)

// GetTransactionHistory https://developer.apple.com/documentation/appstoreserverapi/get_transaction_history
func (this *Client) GetTransactionHistory(transactionId string, param TransactionHistoryParam) (result *TransactionHistoryRsp, err error) {
	err = this.request(http.MethodGet, this.BuildAPI(kTransactionHistory, transactionId), param, nil, &result)
	return result, err
}
