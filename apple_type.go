package apple

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/smartwalle/apple/internal/storekit"
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
	if err := storekit.DecodeClaims(string(s), item); err != nil {
		return nil, err
	}
	return item, nil
}

type InAppOwnershipType string

const (
	InAppOwnershipTypeFamilyShared InAppOwnershipType = "FAMILY_SHARED"
	InAppOwnershipTypePUrchased    InAppOwnershipType = "PURCHASED"
)

type OfferType int

const (
	OfferTypeIntroductory OfferType = 1
	OfferTypePromotional  OfferType = 2
	OfferTypeSubscription OfferType = 3
)

type TransactionType string

const (
	TransactionTypeAutoRenewable TransactionType = "Auto-Renewable Subscription"
	TransactionTypeNonConsumable TransactionType = "Non-Consumable"
	TransactionTypeConsumable    TransactionType = "Consumable"
	TransactionTypeNonRenewing   TransactionType = "Non-Renewing Subscription"
)

// Transaction
// https://developer.apple.com/documentation/appstoreserverapi/jwstransactiondecodedpayload
// https://developer.apple.com/documentation/appstoreservernotifications/jwstransactiondecodedpayload
type Transaction struct {
	jwt.RegisteredClaims
	AppAccountToken             string             `json:"appAccountToken"`
	TransactionId               string             `json:"transactionId"`
	OriginalTransactionId       string             `json:"originalTransactionId"`
	WebOrderLineItemId          string             `json:"webOrderLineItemId"`
	BundleId                    string             `json:"bundleId"`
	ProductId                   string             `json:"productId"`
	SubscriptionGroupIdentifier string             `json:"subscriptionGroupIdentifier"`
	PurchaseDate                int64              `json:"purchaseDate"`
	OriginalPurchaseDate        int64              `json:"originalPurchaseDate"`
	ExpiresDate                 int64              `json:"expiresDate"`
	Quantity                    int                `json:"quantity"`
	Type                        TransactionType    `json:"type"`
	InAppOwnershipType          InAppOwnershipType `json:"inAppOwnershipType"`
	SignedDate                  int64              `json:"signedDate"`
	OfferIdentifier             string             `json:"offerIdentifier"`
	OfferType                   OfferType          `json:"offerType"`
	Environment                 Environment        `json:"environment"`
	RevocationReason            int                `json:"revocationReason"`
	RevocationDate              int64              `json:"revocationDate"`
	IsUpgraded                  bool               `json:"isUpgraded"`
}

type SignedRenewal string

func (s SignedRenewal) Decode() (*Renewal, error) {
	if s == "" {
		return nil, nil
	}
	var item = &Renewal{}
	if err := storekit.DecodeClaims(string(s), item); err != nil {
		return nil, err
	}
	return item, nil
}

type AutoRenewStatus int

const (
	AutoRenewStatusOff AutoRenewStatus = 0
	AutoRenewStatusOn  AutoRenewStatus = 1
)

// Renewal
// https://developer.apple.com/documentation/appstoreserverapi/jwsrenewalinfodecodedpayload
// https://developer.apple.com/documentation/appstoreservernotifications/jwsrenewalinfodecodedpayload
type Renewal struct {
	jwt.RegisteredClaims
	AutoRenewProductId          string          `json:"autoRenewProductId"`
	AutoRenewStatus             AutoRenewStatus `json:"autoRenewStatus"`
	Environment                 Environment     `json:"environment"`
	ExpirationIntent            int             `json:"expirationIntent"`
	GracePeriodExpiresDate      int64           `json:"gracePeriodExpiresDate"`
	IsInBillingRetryPeriod      bool            `json:"isInBillingRetryPeriod"`
	OfferIdentifier             string          `json:"offerIdentifier"`
	OfferType                   OfferType       `json:"offerType"`
	OriginalTransactionId       string          `json:"originalTransactionId"`
	PriceIncreaseStatus         int             `json:"priceIncreaseStatus"`
	ProductId                   string          `json:"productId"`
	RecentSubscriptionStartDate int64           `json:"recentSubscriptionStartDate"`
	SignedDate                  int64           `json:"signedDate"`
}
