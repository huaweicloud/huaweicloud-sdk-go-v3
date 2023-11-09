package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCgwRequest Request Object
type DeleteCgwRequest struct {

	// 对端网关ID
	CustomerGatewayId string `json:"customer_gateway_id"`
}

func (o DeleteCgwRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCgwRequest struct{}"
	}

	return strings.Join([]string{"DeleteCgwRequest", string(data)}, " ")
}
