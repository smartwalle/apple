package apple

import "net/http"

const (
	kRefundLookup = "/v2/refund/lookup/"
)

// RefundLookup https://developer.apple.com/documentation/appstoreserverapi/get_refund_history
func (c *Client) RefundLookup(transactionId string, param RefundLookupParam) (result *RefundLookupResponse, err error) {
	err = c.request(http.MethodGet, c.BuildAPI(kRefundLookup, transactionId), param, nil, &result)
	return result, err
}
