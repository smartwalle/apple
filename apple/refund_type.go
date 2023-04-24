package apple

import (
	"net/url"
)

// RefundLookupParam https://developer.apple.com/documentation/appstoreserverapi/get_refund_history#query-parameters
type RefundLookupParam struct {
	Revision string
}

func (this RefundLookupParam) Values() url.Values {
	var values = url.Values{}
	if this.Revision != "" {
		values.Set("revision", this.Revision)
	}
	return values
}

// RefundLookupRsp https://developer.apple.com/documentation/appstoreserverapi/refundhistoryresponse
type RefundLookupRsp struct {
	HasMore            bool                `json:"hasMore"`
	Revision           string              `json:"revision"`
	SignedTransactions []SignedTransaction `json:"signedTransactions"`
}
