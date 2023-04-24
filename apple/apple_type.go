package apple

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/smartwalle/inpay/apple/internal"
	"net/url"
)

type Environment string

const (
	EnvironmentSandbox    Environment = "Sandbox"
	EnvironmentProduction Environment = "Production"
)

type Param interface {
	Values() url.Values
}

type ResponseError struct {
	Code    int    `json:"errorCode"`
	Message string `json:"errorMessage"`
}

func (this *ResponseError) Error() string {
	return fmt.Sprintf("%d - %s", this.Code, this.Message)
}

type SignedTransaction string

func (s SignedTransaction) Decode() (*Transaction, error) {
	if s == "" {
		return nil, nil
	}
	var item = &Transaction{}
	if err := internal.DecodeClaims(string(s), item); err != nil {
		return nil, err
	}
	return item, nil
}

// Transaction https://developer.apple.com/documentation/appstoreserverapi/jwstransactiondecodedpayload
type Transaction struct {
	jwt.RegisteredClaims
	TransactionId               string      `json:"transactionId"`
	OriginalTransactionId       string      `json:"originalTransactionId"`
	WebOrderLineItemId          string      `json:"webOrderLineItemId"`
	BundleId                    string      `json:"bundleId"`
	ProductId                   string      `json:"productId"`
	SubscriptionGroupIdentifier string      `json:"subscriptionGroupIdentifier"`
	PurchaseDate                int64       `json:"purchaseDate"`
	OriginalPurchaseDate        int64       `json:"originalPurchaseDate"`
	ExpiresDate                 int64       `json:"expiresDate"`
	Quantity                    int         `json:"quantity"`
	Type                        string      `json:"type"`
	InAppOwnershipType          string      `json:"inAppOwnershipType"`
	SignedDate                  int64       `json:"signedDate"`
	OfferType                   int         `json:"offerType"`
	Environment                 Environment `json:"environment"`
	RevocationReason            int         `json:"revocationReason"`
	RevocationDate              int64       `json:"revocationDate"`
}

type SignedRenewal string

func (s SignedRenewal) Decode() (*RenewalInfo, error) {
	if s == "" {
		return nil, nil
	}
	var item = &RenewalInfo{}
	if err := internal.DecodeClaims(string(s), item); err != nil {
		return nil, err
	}
	return item, nil
}

// RenewalInfo https://developer.apple.com/documentation/appstoreserverapi/jwsrenewalinfodecodedpayload
type RenewalInfo struct {
	jwt.RegisteredClaims
	AutoRenewProductId          string      `json:"autoRenewProductId"`
	AutoRenewStatus             int         `json:"autoRenewStatus"`
	Environment                 Environment `json:"environment"`
	ExpirationIntent            int         `json:"expirationIntent"`
	GracePeriodExpiresDate      int64       `json:"gracePeriodExpiresDate"`
	IsInBillingRetryPeriod      bool        `json:"isInBillingRetryPeriod"`
	OfferIdentifier             string      `json:"offerIdentifier"`
	OfferType                   int         `json:"offerType"`
	OriginalTransactionId       string      `json:"originalTransactionId"`
	PriceIncreaseStatus         int         `json:"priceIncreaseStatus"`
	ProductId                   string      `json:"productId"`
	RecentSubscriptionStartDate int64       `json:"recentSubscriptionStartDate"`
	SignedDate                  int64       `json:"signedDate"`
}
