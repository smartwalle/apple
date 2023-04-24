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
