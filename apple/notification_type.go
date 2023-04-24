package apple

// TestNotificationRsp https://developer.apple.com/documentation/appstoreserverapi/sendtestnotificationresponse
type TestNotificationRsp struct {
	TestNotificationToken string `json:"testNotificationToken"`
}
