package apple

// OrderLookupRsp https://developer.apple.com/documentation/appstoreserverapi/orderlookupresponse
type OrderLookupRsp struct {
	Status             int                 `json:"status"`
	SignedTransactions []SignedTransaction `json:"signedTransactions"`
}
