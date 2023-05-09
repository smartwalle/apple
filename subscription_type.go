package apple

import "encoding/json"

// SubscriptionsStatusResponse https://developer.apple.com/documentation/appstoreserverapi/statusresponse
type SubscriptionsStatusResponse struct {
	Data        []*SubscriptionGroupIdentifier `json:"data"`
	Environment Environment                    `json:"environment"`
	AppAppleId  int                            `json:"appAppleId"`
	BundleId    string                         `json:"bundleId"`
}

type SubscriptionGroupIdentifier struct {
	SubscriptionGroupIdentifier string             `json:"subscriptionGroupIdentifier"`
	LastTransactions            []*LastTransaction `json:"lastTransactions"`
}

type LastTransaction struct {
	OriginalTransactionId string       `json:"originalTransactionId"`
	Status                int          `json:"status"`
	Renewal               *Renewal     `json:"renewal"`
	Transaction           *Transaction `json:"transaction"`
}

type LastTransactionAlias LastTransaction

func (this *LastTransaction) UnmarshalJSON(data []byte) (err error) {
	var aux = struct {
		*LastTransactionAlias
		SignedRenewal     SignedRenewal     `json:"signedRenewalInfo"`
		SignedTransaction SignedTransaction `json:"signedTransactionInfo"`
	}{
		LastTransactionAlias: (*LastTransactionAlias)(this),
	}

	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if this.Renewal, err = aux.SignedRenewal.Decode(); err != nil {
		return err
	}
	if this.Transaction, err = aux.SignedTransaction.Decode(); err != nil {
		return err
	}
	return nil
}

// ExtendRenewalDateParam https://developer.apple.com/documentation/appstoreserverapi/extendrenewaldaterequest
type ExtendRenewalDateParam struct {
	ExtendByDays      int    `json:"extendByDays"`
	ExtendReasonCode  int    `json:"extendReasonCode"`
	RequestIdentifier string `json:"requestIdentifier"`
}

// ExtendRenewalDateResponse https://developer.apple.com/documentation/appstoreserverapi/extendrenewaldateresponse
type ExtendRenewalDateResponse struct {
	EffectiveDate         int64  `json:"effectiveDate"`
	OriginalTransactionId string `json:"originalTransactionId"`
	Success               bool   `json:"success"`
	WebOrderLineItemId    string `json:"webOrderLineItemId"`
}
