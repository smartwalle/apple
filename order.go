package apple

import "net/http"

const (
	kOrderLookup = "/v1/lookup/"
)

// OrderLookup https://developer.apple.com/documentation/appstoreserverapi/look_up_order_id
func (c *Client) OrderLookup(orderId string) (result *OrderLookupResponse, err error) {
	err = c.request(http.MethodGet, c.BuildAPI(kOrderLookup, orderId), nil, nil, &result)
	return result, err
}
