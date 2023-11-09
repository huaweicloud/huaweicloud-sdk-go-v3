package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCgwResponse Response Object
type ShowCgwResponse struct {
	CustomerGateway *ResponseCustomerGateway `json:"customer_gateway,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCgwResponse struct{}"
	}

	return strings.Join([]string{"ShowCgwResponse", string(data)}, " ")
}
