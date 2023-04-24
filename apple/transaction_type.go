package apple

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/smartwalle/inpay/apple/internal"
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

// TransactionHistoryRsp https://developer.apple.com/documentation/appstoreserverapi/historyresponse
type TransactionHistoryRsp struct {
	AppAppleId         int                 `json:"appAppleId"`
	BundleId           string              `json:"bundleId"`
	Environment        Environment         `json:"environment"`
	HasMore            bool                `json:"hasMore"`
	Revision           string              `json:"revision"`
	SignedTransactions []SignedTransaction `json:"signedTransactions"`
}

type SignedTransaction string

func (s SignedTransaction) DecodeSignedTransaction() (*TransactionItem, error) {
	if s == "" {
		return nil, nil
	}
	var item = &TransactionItem{}
	if err := internal.DecodeClaims(string(s), item); err != nil {
		return nil, err
	}
	return item, nil
}

type TransactionItem struct {
	jwt.RegisteredClaims
	TransactionId               string `json:"transactionId"`
	OriginalTransactionId       string `json:"originalTransactionId"`
	WebOrderLineItemId          string `json:"webOrderLineItemId"`
	BundleId                    string `json:"bundleId"`
	ProductId                   string `json:"productId"`
	SubscriptionGroupIdentifier string `json:"subscriptionGroupIdentifier"`
	PurchaseDate                int64  `json:"purchaseDate"`
	OriginalPurchaseDate        int64  `json:"originalPurchaseDate"`
	ExpiresDate                 int64  `json:"expiresDate"`
	Quantity                    int    `json:"quantity"`
	Type                        string `json:"type"`
	InAppOwnershipType          string `json:"inAppOwnershipType"`
	SignedDate                  int64  `json:"signedDate"`
	OfferType                   int    `json:"offerType"`
	Environment                 string `json:"environment"`
	RevocationReason            int    `json:"revocationReason"`
	RevocationDate              int64  `json:"revocationDate"`
}
