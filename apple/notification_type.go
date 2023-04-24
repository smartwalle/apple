package apple

import "github.com/golang-jwt/jwt/v5"

// TestNotificationRsp https://developer.apple.com/documentation/appstoreserverapi/sendtestnotificationresponse
type TestNotificationRsp struct {
	TestNotificationToken string `json:"testNotificationToken"`
}

// Notification https://developer.apple.com/documentation/appstoreservernotifications/responsebodyv2decodedpayload
type Notification struct {
	jwt.RegisteredClaims
	NotificationType string               `json:"notificationType"`
	Subtype          string               `json:"subtype"`
	Data             *NotificationData    `json:"data"`
	Summary          *NotificationSummary `json:"summary"`
	Version          string               `json:"version"`
	SignedDate       int64                `json:"signedDate"`
	NotificationUUID string               `json:"notificationUUID"`
}

type NotificationData struct {
	AppAppleId         int64             `json:"appAppleId"`
	BundleId           string            `json:"bundleId"`
	BundleVersion      string            `json:"bundleVersion"`
	Environment        Environment       `json:"environment"`
	SignedRenewalInfo  SignedRenewal     `json:"signedRenewalInfo"`
	SignedTransactions SignedTransaction `json:"signedTransactionInfo"`
}

type NotificationSummary struct {
	RequestIdentifier      string      `json:"requestIdentifier"`
	Environment            Environment `json:"environment"`
	AppAppleId             int64       `json:"appAppleId"`
	BundleId               string      `json:"bundleId"`
	ProductId              string      `json:"productId"`
	StorefrontCountryCodes []string    `json:"storefrontCountryCodes"`
	FailedCount            int64       `json:"failedCount"`
	SucceededCount         int64       `json:"succeededCount"`
}
