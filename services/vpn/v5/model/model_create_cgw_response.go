package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCgwResponse Response Object
type CreateCgwResponse struct {
	CustomerGateway *ResponseCustomerGateway `json:"customer_gateway,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CreateCgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCgwResponse struct{}"
	}

	return strings.Join([]string{"CreateCgwResponse", string(data)}, " ")
}
