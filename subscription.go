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
func (c *Client) GetSubscriptionsStatuses(transactionId string) (result *SubscriptionsStatusResponse, err error) {
	err = c.request(http.MethodGet, c.BuildAPI(kGetSubscriptions, transactionId), nil, nil, &result)
	return result, err
}

// ExtendSubscription https://developer.apple.com/documentation/appstoreserverapi/extend_a_subscription_renewal_date
func (c *Client) ExtendSubscription(transactionId string, param ExtendRenewalDateParam) (result *ExtendRenewalDateResponse, err error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	err = c.request(http.MethodPut, c.BuildAPI(kExtendSubscription, transactionId), nil, bytes.NewReader(data), &result)
	return result, err
}
