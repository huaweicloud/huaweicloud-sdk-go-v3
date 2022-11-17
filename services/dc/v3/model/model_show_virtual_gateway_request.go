package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowVirtualGatewayRequest struct {

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// 虚拟网关ID
	VirtualGatewayId string `json:"virtual_gateway_id"`
}

func (o ShowVirtualGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVirtualGatewayRequest struct{}"
	}

	return strings.Join([]string{"ShowVirtualGatewayRequest", string(data)}, " ")
}
