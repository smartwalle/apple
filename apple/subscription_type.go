package apple

// SubscriptionsStatusRsp https://developer.apple.com/documentation/appstoreserverapi/statusresponse
type SubscriptionsStatusRsp struct {
	Data        []*SubscriptionGroupIdentifierItem `json:"data"`
	Environment Environment                        `json:"environment"`
	AppAppleId  int                                `json:"appAppleId"`
	BundleId    string                             `json:"bundleId"`
}

type SubscriptionGroupIdentifierItem struct {
	SubscriptionGroupIdentifier string                  `json:"subscriptionGroupIdentifier"`
	LastTransactions            []*LastTransactionsItem `json:"lastTransactions"`
}

type LastTransactionsItem struct {
	OriginalTransactionId string            `json:"originalTransactionId"`
	Status                int               `json:"status"`
	SignedRenewalInfo     SignedRenewal     `json:"signedRenewalInfo"`
	SignedTransactionInfo SignedTransaction `json:"signedTransactionInfo"`
}

// ExtendRenewalDateParam https://developer.apple.com/documentation/appstoreserverapi/extendrenewaldaterequest
type ExtendRenewalDateParam struct {
	ExtendByDays      int    `json:"extendByDays"`
	ExtendReasonCode  int    `json:"extendReasonCode"`
	RequestIdentifier string `json:"requestIdentifier"`
}

// ExtendRenewalDateRsp https://developer.apple.com/documentation/appstoreserverapi/extendrenewaldateresponse
type ExtendRenewalDateRsp struct {
	EffectiveDate         int64  `json:"effectiveDate"`
	OriginalTransactionId string `json:"originalTransactionId"`
	Success               bool   `json:"success"`
	WebOrderLineItemId    string `json:"webOrderLineItemId"`
}
