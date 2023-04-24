package apple

import "net/http"

const (
	kGetSubscriptions = "/v1/subscriptions/"
)

// GetSubscriptionsStatuses https://developer.apple.com/documentation/appstoreserverapi/get_all_subscription_statuses
func (this *Client) GetSubscriptionsStatuses(transactionId string) (result *SubscriptionsStatusRsp, err error) {
	err = this.request(http.MethodGet, this.BuildAPI(kGetSubscriptions, transactionId), nil, &result)
	return result, err
}
