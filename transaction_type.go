package apple

import (
	"encoding/json"
	"fmt"
	"net/url"
)

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
	ExcludeRevoked              bool
}

func (this TransactionHistoryParam) Values() url.Values {
	var values = url.Values{}
	if this.Revision != "" {
		values.Set("revision", this.Revision)
	}
	if this.StartDate != "" {
		values.Set("startDate", this.StartDate)
	}
	if this.EndDate != "" {
		values.Set("endDate", this.EndDate)
	}
	if this.ProductId != "" {
		values.Set("productId", this.ProductId)
	}
	if this.ProductType != "" {
		values.Set("productType", this.ProductType)
	}
	if this.Sort != "" {
		values.Set("sort", this.Sort)
	}
	if this.SubscriptionGroupIdentifier != "" {
		values.Set("subscriptionGroupIdentifier", this.SubscriptionGroupIdentifier)
	}
	if this.InAppOwnershipType != "" {
		values.Set("inAppOwnershipType", this.InAppOwnershipType)
	}
	values.Set("excludeRevoked", fmt.Sprintf("%v", this.ExcludeRevoked))
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

func (this *TransactionHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	var aux = struct {
		*TransactionHistoryResponseAlias
		SignedTransactions []SignedTransaction `json:"signedTransactions"`
	}{
		TransactionHistoryResponseAlias: (*TransactionHistoryResponseAlias)(this),
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
