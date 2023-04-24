package apple

// 21000 App Store 无法读取你提供的JSON数据
// 21002 收据数据不符合格式
// 21003 收据无法被验证
// 21004 你提供的共享密钥和账户的共享密钥不一致
// 21005 收据服务器当前不可用
// 21006 收据是有效的，但订阅服务已经过期。当收到这个信息时，解码后的收据信息也包含在返回内容中
// 21007 收据信息是测试用（sandbox），但却被发送到产品环境中验证
// 21008 收据信息是产品环境中使用，但却被发送到测试环境中验证

type Environment string

const (
	EnvironmentSandbox    Environment = "Sandbox"
	EnvironmentProduction Environment = "Production"
)

type GetReceiptParam struct {
	Receipt  string `json:"receipt-data"`
	Password string `json:"password,omitempty"`
}

type ReceiptSummary struct {
	Receipt     *Receipt    `json:"receipt"`
	Environment Environment `json:"environment"`
	Status      int         `json:"status"`
}

type Receipt struct {
	ReceiptType                string   `json:"receipt_type"`
	AdamId                     int64    `json:"adam_id"`
	AppItemId                  int64    `json:"app_item_id"`
	BundleId                   string   `json:"bundle_id"`
	ApplicationVersion         string   `json:"application_version"`
	DownloadId                 int64    `json:"download_id"`
	VersionExternalIdentifier  int64    `json:"version_external_identifier"`
	ReceiptCreationDate        string   `json:"receipt_creation_date"`
	ReceiptCreationDateMS      string   `json:"receipt_creation_date_ms"`
	ReceiptCreationDatePst     string   `json:"receipt_creation_date_pst"`
	RequestDate                string   `json:"request_date"`
	RequestDateMS              string   `json:"request_date_ms"`
	RequestDatePst             string   `json:"request_date_pst"`
	OriginalPurchaseDate       string   `json:"original_purchase_date"`
	OriginalPurchaseDateMS     string   `json:"original_purchase_date_ms"`
	OriginalPurchaseDatePst    string   `json:"original_purchase_date_pst"`
	OriginalApplicationVersion string   `json:"original_application_version"`
	InApp                      []*InApp `json:"in_app"`
}

type InApp struct {
	Quantity                string `json:"quantity"`
	ProductId               string `json:"product_id"`
	TransactionId           string `json:"transaction_id"`
	OriginalTransactionId   string `json:"original_transaction_id"`
	PurchaseDate            string `json:"purchase_date"`
	PurchaseDateMS          string `json:"purchase_date_ms"`
	PurchaseDatePst         string `json:"purchase_date_pst"`
	OriginalPurchaseDate    string `json:"original_purchase_date"`
	OriginalPurchaseDateMS  string `json:"original_purchase_date_ms"`
	OriginalPurchaseDatePst string `json:"original_purchase_date_pst"`
	IsTrialPeriod           string `json:"is_trial_period"`
	InAppOwnershipType      string `json:"in_app_ownership_type"`
}
