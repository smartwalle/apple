package apple

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// TransactionResponse https://developer.apple.com/documentation/appstoreserverapi/transactioninforesponse
type TransactionResponse struct {
	*Transaction `json:"transaction"`
}

type TransactionResponseAlias TransactionResponse

func (t *TransactionResponse) UnmarshalJSON(data []byte) (err error) {
	var aux = struct {
		*TransactionResponseAlias
		SignedTransactions SignedTransaction `json:"signedTransactionInfo"`
	}{
		TransactionResponseAlias: (*TransactionResponseAlias)(t),
	}

	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var transaction *Transaction
	transaction, err = aux.SignedTransactions.Decode()
	if err != nil {
		return err
	}
	t.Transaction = transaction
	return nil
}

// TransactionHistoryParam https://developer.apple.com/documentation/appstoreserverapi/get_transaction_history#query-parameters
type TransactionHistoryParam struct {
	Revision                    string
	StartDate                   string
	EndDate                     string
	ProductId                   string
	ProductType                 string
	Sort                        string
	SubscriptionGroupIdentifier string
	InAppOwnershipType          string
	Revoked                     bool
}

func (t TransactionHistoryParam) Values() url.Values {
	var values = url.Values{}
	if t.Revision != "" {
		values.Set("revision", t.Revision)
	}
	if t.StartDate != "" {
		values.Set("startDate", t.StartDate)
	}
	if t.EndDate != "" {
		values.Set("endDate", t.EndDate)
	}
	if t.ProductId != "" {
		values.Set("productId", t.ProductId)
	}
	if t.ProductType != "" {
		values.Set("productType", t.ProductType)
	}
	if t.Sort != "" {
		values.Set("sort", t.Sort)
	}
	if t.SubscriptionGroupIdentifier != "" {
		values.Set("subscriptionGroupIdentifier", t.SubscriptionGroupIdentifier)
	}
	if t.InAppOwnershipType != "" {
		values.Set("inAppOwnershipType", t.InAppOwnershipType)
	}
	values.Set("revoked", fmt.Sprintf("%v", t.Revoked))
	return values
}

// TransactionHistoryResponse https://developer.apple.com/documentation/appstoreserverapi/historyresponse
type TransactionHistoryResponse struct {
	AppAppleId   int            `json:"appAppleId"`
	BundleId     string         `json:"bundleId"`
	Environment  Environment    `json:"environment"`
	HasMore      bool           `json:"hasMore"`
	Revision     string         `json:"revision"`
	Transactions []*Transaction `json:"transactions"`
}

type TransactionHistoryResponseAlias TransactionHistoryResponse

func (t *TransactionHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	var aux = struct {
		*TransactionHistoryResponseAlias
		SignedTransactions []SignedTransaction `json:"signedTransactions"`
	}{
		TransactionHistoryResponseAlias: (*TransactionHistoryResponseAlias)(t),
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
		t.Transactions = append(t.Transactions, transaction)
	}
	return nil
}

// ConsumptionParam https://developer.apple.com/documentation/appstoreserverapi/consumptionrequest
type ConsumptionParam struct {
	AccountTenure            int    `json:"accountTenure"`
	AppAccountToken          string `json:"appAccountToken"`
	ConsumptionStatus        int    `json:"consumptionStatus"`
	CustomerConsented        bool   `json:"customerConsented"`
	DeliveryStatus           int    `json:"deliveryStatus"`
	LifetimeDollarsPurchased int    `json:"lifetimeDollarsPurchased"`
	LifetimeDollarsRefunded  int    `json:"lifetimeDollarsRefunded"`
	Platform                 int    `json:"platform"`
	PlayTime                 int    `json:"playTime"`
	SampleContentProvided    bool   `json:"sampleContentProvided"`
	UserStatus               int    `json:"userStatus"`
}
