package apple

import "net/http"

// 21000 App Store 无法读取你提供的JSON数据
// 21002 收据数据不符合格式
// 21003 收据无法被验证
// 21004 你提供的共享密钥和账户的共享密钥不一致
// 21005 收据服务器当前不可用
// 21006 收据是有效的，但订阅服务已经过期。当收到这个信息时，解码后的收据信息也包含在返回内容中
// 21007 收据信息是测试用（sandbox），但却被发送到产品环境中验证
// 21008 收据信息是产品环境中使用，但却被发送到测试环境中验证

type ReceiptOptions struct {
	Client                 *http.Client `json:"-"`
	Receipt                string       `json:"receipt-data"`
	Password               string       `json:"password,omitempty"`
	ExcludeOldTransactions bool         `json:"exclude-old-transactions"`
}

type VerifyReceiptOptionFunc func(opts *ReceiptOptions)

func WithHTTPClient(client *http.Client) VerifyReceiptOptionFunc {
	return func(opts *ReceiptOptions) {
		if client != nil {
			opts.Client = client
		}
	}
}

func WithPassword(password string) VerifyReceiptOptionFunc {
	return func(opts *ReceiptOptions) {
		opts.Password = password
	}
}

func WithExcludeOldTransactions(value bool) VerifyReceiptOptionFunc {
	return func(opts *ReceiptOptions) {
		opts.ExcludeOldTransactions = value
	}
}

type ReceiptSummary struct {
	Environment        Environment           `json:"environment"`
	IsRetryable        bool                  `json:"is_retryable"`
	LatestReceipt      string                `json:"latest_receipt,omitempty"`
	LatestReceiptInfo  []*LatestReceiptInfo  `json:"latest_receipt_info,omitempty"`
	PendingRenewalInfo []*PendingRenewalInfo `json:"pending_renewal_info,omitempty"`
	Receipt            *Receipt              `json:"receipt"`
	Status             int                   `json:"status"`
}

type Receipt struct {
	AdamId                     int64    `json:"adam_id"`
	AppItemId                  int64    `json:"app_item_id"`
	ApplicationVersion         string   `json:"application_version"`
	BundleId                   string   `json:"bundle_id"`
	DownloadId                 int64    `json:"download_id"`
	ExpirationDate             string   `json:"expiration_date"`
	ExpirationDateMs           string   `json:"expiration_date_ms"`
	ExpirationDatePST          string   `json:"expiration_date_pst"`
	OriginalPurchaseDate       string   `json:"original_purchase_date"`
	OriginalPurchaseDateMS     string   `json:"original_purchase_date_ms"`
	OriginalPurchaseDatePST    string   `json:"original_purchase_date_pst"`
	OriginalApplicationVersion string   `json:"original_application_version"`
	PreorderDate               string   `json:"preorder_date"`
	PreorderDateMS             string   `json:"preorder_date_ms"`
	PreorderDatePST            string   `json:"preorder_date_pst"`
	ReceiptCreationDate        string   `json:"receipt_creation_date"`
	ReceiptCreationDateMS      string   `json:"receipt_creation_date_ms"`
	ReceiptCreationDatePST     string   `json:"receipt_creation_date_pst"`
	ReceiptType                string   `json:"receipt_type"`
	RequestDate                string   `json:"request_date"`
	RequestDateMS              string   `json:"request_date_ms"`
	RequestDatePST             string   `json:"request_date_pst"`
	VersionExternalIdentifier  int64    `json:"version_external_identifier"`
	InApp                      []*InApp `json:"in_app"`
}

type InApp struct {
	CancellationDate        string `json:"cancellation_date"`
	CancellationDateMs      string `json:"cancellation_date_ms"`
	CancellationDatePST     string `json:"cancellation_date_pst"`
	CancellationReason      string `json:"cancellation_reason"`
	ExpiresDate             string `json:"expires_date"`
	ExpiresDateMs           string `json:"expires_date_ms"`
	ExpiresDatePST          string `json:"expires_date_pst"`
	IsInIntroOfferPeriod    string `json:"is_in_intro_offer_period"`
	IsTrialPeriod           string `json:"is_trial_period"`
	OriginalPurchaseDate    string `json:"original_purchase_date"`
	OriginalPurchaseDateMS  string `json:"original_purchase_date_ms"`
	OriginalPurchaseDatePST string `json:"original_purchase_date_pst"`
	OriginalTransactionId   string `json:"original_transaction_id"`
	ProductId               string `json:"product_id"`
	PromotionalOfferId      string `json:"promotional_offer_id"`
	PurchaseDate            string `json:"purchase_date"`
	PurchaseDateMS          string `json:"purchase_date_ms"`
	PurchaseDatePST         string `json:"purchase_date_pst"`
	Quantity                string `json:"quantity"`
	TransactionId           string `json:"transaction_id"`
	WebOrderLineItemId      string `json:"web_order_line_item_id"`
}

type LatestReceiptInfo struct {
	AppAccountToken             string `json:"app_account_token"`
	CancellationDate            string `json:"cancellation_date"`
	CancellationDateMs          string `json:"cancellation_date_ms"`
	CancellationDatePST         string `json:"cancellation_date_pst"`
	CancellationReason          string `json:"cancellation_reason"`
	ExpiresDate                 string `json:"expires_date"`
	ExpiresDateMs               string `json:"expires_date_ms"`
	ExpiresDatePST              string `json:"expires_date_pst"`
	InAppOwnershipType          string `json:"in_app_ownership_type"`
	IsInIntroOfferPeriod        string `json:"is_in_intro_offer_period"`
	IsTrialPeriod               string `json:"is_trial_period"`
	IsUpgraded                  string `json:"is_upgraded"`
	OfferCodeRefName            string `json:"offer_code_ref_name"`
	OriginalPurchaseDate        string `json:"original_purchase_date"`
	OriginalPurchaseDateMS      string `json:"original_purchase_date_ms"`
	OriginalPurchaseDatePST     string `json:"original_purchase_date_pst"`
	OriginalTransactionId       string `json:"original_transaction_id"`
	ProductId                   string `json:"product_id"`
	PromotionalOfferId          string `json:"promotional_offer_id"`
	PurchaseDate                string `json:"purchase_date"`
	PurchaseDateMS              string `json:"purchase_date_ms"`
	PurchaseDatePST             string `json:"purchase_date_pst"`
	Quantity                    string `json:"quantity"`
	SubscriptionGroupIdentifier string `json:"subscription_group_identifier"`
	TransactionId               string `json:"transaction_id"`
	WebOrderLineItemId          string `json:"web_order_line_item_id"`
}

type PendingRenewalInfo struct {
	SubscriptionExpirationIntent   string `json:"expiration_intent"`
	SubscriptionAutoRenewProductID string `json:"auto_renew_product_id"`
	SubscriptionRetryFlag          string `json:"is_in_billing_retry_period"`
	SubscriptionAutoRenewStatus    string `json:"auto_renew_status"`
	SubscriptionPriceConsentStatus string `json:"price_consent_status"`
	ProductID                      string `json:"product_id"`
	OriginalTransactionID          string `json:"original_transaction_id"`
	OfferCodeRefName               string `json:"offer_code_ref_name,omitempty"`
	PromotionalOfferID             string `json:"promotional_offer_id,omitempty"`
	PriceIncreaseStatus            string `json:"price_increase_status,omitempty"`
	GracePeriodDate                string `json:"grace_period_expires_date,omitempty"`
	GracePeriodDateMS              string `json:"grace_period_expires_date_ms,omitempty"`
	GracePeriodDatePST             string `json:"grace_period_expires_date_pst,omitempty"`
}
