package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCgwResponse Response Object
type UpdateCgwResponse struct {
	CustomerGateway *ResponseCustomerGateway `json:"customer_gateway,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCgwResponse struct{}"
	}

	return strings.Join([]string{"UpdateCgwResponse", string(data)}, " ")
}
