package apple

import "net/http"

const (
	kOrderLookup = "/lookup/"
)

// OrderLookup https://developer.apple.com/documentation/appstoreserverapi/look_up_order_id
func (this *Client) OrderLookup(orderId string) (result *OrderLookupRsp, err error) {
	err = this.request(http.MethodGet, this.BuildAPI(kOrderLookup, orderId), nil, &result)
	return result, err
}
