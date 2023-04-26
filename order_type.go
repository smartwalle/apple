package apple

import "encoding/json"

// OrderLookupResponse https://developer.apple.com/documentation/appstoreserverapi/orderlookupresponse
type OrderLookupResponse struct {
	Status       int            `json:"status"`
	Transactions []*Transaction `json:"transactions"`
}

type OrderLookupResponseAlias OrderLookupResponse

func (this *OrderLookupResponse) UnmarshalJSON(data []byte) (err error) {
	var aux = &struct {
		*OrderLookupResponseAlias
		SignedTransactions []SignedTransaction `json:"signedTransactions"`
	}{
		OrderLookupResponseAlias: (*OrderLookupResponseAlias)(this),
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
		this.Transactions = append(this.Transactions, transaction)
	}
	return nil
}
