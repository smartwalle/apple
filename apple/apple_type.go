package apple

import (
	"fmt"
	"net/url"
)

type Param interface {
	Values() url.Values
}

type ResponseError struct {
	Code    int    `json:"errorCode"`
	Message string `json:"errorMessage"`
}

func (this *ResponseError) Error() string {
	return fmt.Sprintf("%d - %s", this.Code, this.Message)
}
