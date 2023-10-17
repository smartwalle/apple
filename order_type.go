package apple

import "encoding/json"

// OrderLookupResponse https://developer.apple.com/documentation/appstoreserverapi/orderlookupresponse
type OrderLookupResponse struct {
	Status       int            `json:"status"`
	Transactions []*Transaction `json:"transactions"`
}

type OrderLookupResponseAlias OrderLookupResponse

func (o *OrderLookupResponse) UnmarshalJSON(data []byte) (err error) {
	var aux = struct {
		*OrderLookupResponseAlias
		SignedTransactions []SignedTransaction `json:"signedTransactions"`
	}{
		OrderLookupResponseAlias: (*OrderLookupResponseAlias)(o),
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
		o.Transactions = append(o.Transactions, transaction)
	}
	return nil
}
