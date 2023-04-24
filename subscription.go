package apple

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	kGetSubscriptions   = "/v1/subscriptions/"
	kExtendSubscription = "/v1/subscriptions/extend/"
)

// GetSubscriptionsStatuses https://developer.apple.com/documentation/appstoreserverapi/get_all_subscription_statuses
func (this *Client) GetSubscriptionsStatuses(transactionId string) (result *SubscriptionsStatusRsp, err error) {
	err = this.request(http.MethodGet, this.BuildAPI(kGetSubscriptions, transactionId), nil, nil, &result)
	return result, err
}

// ExtendSubscription https://developer.apple.com/documentation/appstoreserverapi/extend_a_subscription_renewal_date
func (this *Client) ExtendSubscription(transactionId string, param ExtendRenewalDateParam) (result *ExtendRenewalDateRsp, err error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	err = this.request(http.MethodPut, this.BuildAPI(kExtendSubscription, transactionId), nil, bytes.NewReader(data), &result)
	return result, err
}
