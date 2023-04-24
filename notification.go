package apple

import (
	"encoding/json"
	"github.com/smartwalle/apple/internal"
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

// DecodeNotification 用于解析通知数据 https://developer.apple.com/documentation/appstoreservernotifications/responsebodyv2
//
// 关于接收到苹果服务器推送的通知之后，业务服务器如何响应参照：
// https://developer.apple.com/documentation/appstoreservernotifications/responding_to_app_store_server_notifications
func (this *Client) DecodeNotification(data []byte) (*Notification, error) {
	return DecodeNotification(data)
}

type NotificationPayload struct {
	SignedPayload string `json:"signedPayload"`
}

func DecodeNotification(data []byte) (*Notification, error) {
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
