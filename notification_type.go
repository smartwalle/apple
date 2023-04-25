package apple

import "github.com/golang-jwt/jwt/v5"

// TestNotificationRsp https://developer.apple.com/documentation/appstoreserverapi/sendtestnotificationresponse
type TestNotificationRsp struct {
	TestNotificationToken string `json:"testNotificationToken"`
}

type NotificationType string

const (
	NotificationTypeConsumptionRequest     NotificationType = "CONSUMPTION_REQUEST"
	NotificationTypeDidChangeRenewalPref   NotificationType = "DID_CHANGE_RENEWAL_PREF"
	NotificationTypeDidChangeRenewalStatus NotificationType = "DID_CHANGE_RENEWAL_STATUS"
	NotificationTypeDidFailToRenew         NotificationType = "DID_FAIL_TO_RENEW"
	NotificationTypeDidRenew               NotificationType = "DID_RENEW"
	NotificationTypeExpired                NotificationType = "EXPIRED"
	NotificationTypeGracePeriodExpired     NotificationType = "GRACE_PERIOD_EXPIRED"
	NotificationTypeOfferRedeemed          NotificationType = "OFFER_REDEEMED"
	NotificationTypePriceIncrease          NotificationType = "PRICE_INCREASE"
	NotificationTypeRefund                 NotificationType = "REFUND"
	NotificationTypeRefundDeclined         NotificationType = "REFUND_DECLINED"
	NotificationTypeRenewalExtended        NotificationType = "RENEWAL_EXTENDED"
	NotificationTypeRenewalExtension       NotificationType = "RENEWAL_EXTENSION"
	NotificationTypeRevoke                 NotificationType = "REVOKE"
	NotificationTypeSubscribed             NotificationType = "SUBSCRIBED"
	NotificationTypeTest                   NotificationType = "TEST"
)

type NotificationSubType string

const (
	NotificationSubTypeAccept            NotificationSubType = "ACCEPTED"
	NotificationSubTypeAutoRenewDisabled NotificationSubType = "AUTO_RENEW_DISABLED"
	NotificationSubTypeAutoRenewEnabled  NotificationSubType = "AUTO_RENEW_ENABLED"
	NotificationSubTypeBillingRecovery   NotificationSubType = "BILLING_RECOVERY"
	NotificationSubTypeBillingRetry      NotificationSubType = "BILLING_RETRY"
	NotificationSubTypeDowngrade         NotificationSubType = "DOWNGRADE"
	NotificationSubTypeFailure           NotificationSubType = "FAILURE"
	NotificationSubTypeGracePeriod       NotificationSubType = "GRACE_PERIOD"
	NotificationSubTypeInitialBuy        NotificationSubType = "INITIAL_BUY"
	NotificationSubTypePending           NotificationSubType = "PENDING"
	NotificationSubTypePriceIncrease     NotificationSubType = "PRICE_INCREASE"
	NotificationSubTypeProductNotForSale NotificationSubType = "PRODUCT_NOT_FOR_SALE"
	NotificationSubTypeResubscribe       NotificationSubType = "RESUBSCRIBE"
	NotificationSubTypeSummary           NotificationSubType = "SUMMARY"
	NotificationSubTypeUpgrade           NotificationSubType = "UPGRADE"
	NotificationSubTypeVoluntary         NotificationSubType = "VOLUNTARY"
)

// Notification https://developer.apple.com/documentation/appstoreservernotifications/responsebodyv2decodedpayload
type Notification struct {
	jwt.RegisteredClaims
	NotificationType NotificationType     `json:"notificationType"`
	Subtype          NotificationSubType  `json:"subtype"`
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
