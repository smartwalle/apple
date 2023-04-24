package apple

import (
	"encoding/json"
	"github.com/smartwalle/inpay/apple/internal"
	"net/http"
)

const (
	kTestNotification = "/v1/notifications/test"
)

// RequestTestNotification https://developer.apple.com/documentation/appstoreserverapi/request_a_test_notification
func (this *Client) RequestTestNotification() (result *TestNotificationRsp, err error) {
	err = this.request(http.MethodPost, this.BuildAPI(kTestNotification), nil, nil, &result)
	return result, err
}

// UnmarshalNotification https://developer.apple.com/documentation/appstoreservernotifications/responsebodyv2
// 用于解析通知数据
func (this *Client) UnmarshalNotification(data []byte) (*Notification, error) {
	return UnmarshalNotification(data)
}

type NotificationPayload struct {
	SignedPayload string `json:"signedPayload"`
}

func UnmarshalNotification(data []byte) (*Notification, error) {
	var payload *NotificationPayload
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, err
	}

	var notification = &Notification{}
	if err := internal.DecodeClaims(payload.SignedPayload, notification); err != nil {
		return nil, err
	}
	return notification, nil
}
