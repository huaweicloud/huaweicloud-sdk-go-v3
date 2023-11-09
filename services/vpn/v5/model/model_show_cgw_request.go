package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCgwRequest Request Object
type ShowCgwRequest struct {

	// 对端网关ID
	CustomerGatewayId string `json:"customer_gateway_id"`
}

func (o ShowCgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCgwRequest struct{}"
	}

	return strings.Join([]string{"ShowCgwRequest", string(data)}, " ")
}
