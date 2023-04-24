package apple

import "net/http"

const (
	kTestNotification = "/v1/notifications/test"
)

// RequestTestNotification https://developer.apple.com/documentation/appstoreserverapi/request_a_test_notification
func (this *Client) RequestTestNotification() (result *TestNotificationRsp, err error) {
	err = this.request(http.MethodPost, this.BuildAPI(kTestNotification), nil, nil, &result)
	return result, err
}
