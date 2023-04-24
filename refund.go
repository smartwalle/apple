package apple

import "net/http"

const (
	kRefundLookup = "/v2/refund/lookup/"
)

// RefundLookup https://developer.apple.com/documentation/appstoreserverapi/get_refund_history
func (this *Client) RefundLookup(transactionId string, param RefundLookupParam) (result *RefundLookupRsp, err error) {
	err = this.request(http.MethodGet, this.BuildAPI(kRefundLookup, transactionId), param, nil, &result)
	return result, err
}
