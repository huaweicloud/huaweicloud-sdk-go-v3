package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCgwRequest Request Object
type UpdateCgwRequest struct {

	// 对端网关ID
	CustomerGatewayId string `json:"customer_gateway_id"`

	Body *UpdateCgwRequestBody `json:"body,omitempty"`
}

func (o UpdateCgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCgwRequest struct{}"
	}

	return strings.Join([]string{"UpdateCgwRequest", string(data)}, " ")
}
