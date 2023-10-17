package apple

import (
	"encoding/json"
	"net/url"
)

// RefundLookupParam https://developer.apple.com/documentation/appstoreserverapi/get_refund_history#query-parameters
type RefundLookupParam struct {
	Revision string
}

func (r RefundLookupParam) Values() url.Values {
	var values = url.Values{}
	if r.Revision != "" {
		values.Set("revision", r.Revision)
	}
	return values
}

// RefundLookupResponse https://developer.apple.com/documentation/appstoreserverapi/refundhistoryresponse
type RefundLookupResponse struct {
	HasMore      bool           `json:"hasMore"`
	Revision     string         `json:"revision"`
	Transactions []*Transaction `json:"transactions"`
}

type RefundLookupResponseAlias RefundLookupResponse

func (r *RefundLookupResponse) UnmarshalJSON(data []byte) (err error) {
	var aux = struct {
		*RefundLookupResponseAlias
		SignedTransactions []SignedTransaction `json:"signedTransactions"`
	}{
		RefundLookupResponseAlias: (*RefundLookupResponseAlias)(r),
	}

	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}

	for _, item := range aux.SignedTransactions {
		var transaction *Transaction
		transaction, err = item.Decode()
		if err != nil {
			return err
		}
		r.Transactions = append(r.Transactions, transaction)
	}

	return nil
}
