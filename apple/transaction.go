package apple

import (
	"github.com/smartwalle/ngx"
	"net/http"
)

const (
	kTransactionHistory = "/history/"
)

// GetTransactionHistory https://developer.apple.com/documentation/appstoreserverapi/historyresponse
func (this *Client) GetTransactionHistory(transactionId, revision string) (result *TransactionHistoryRsp, err error) {
	var req = ngx.NewRequest(http.MethodGet, this.BuildAPI(kTransactionHistory, transactionId))
	if revision != "" {
		req.AddQuery("revision", revision)
	}
	err = this.request(req, &result)
	return result, err
}
