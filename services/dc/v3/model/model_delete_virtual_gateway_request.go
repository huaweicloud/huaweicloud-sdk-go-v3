package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVirtualGatewayRequest Request Object
type DeleteVirtualGatewayRequest struct {

	// 虚拟网关ID
	VirtualGatewayId string `json:"virtual_gateway_id"`
}

func (o DeleteVirtualGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVirtualGatewayRequest struct{}"
	}

	return strings.Join([]string{"DeleteVirtualGatewayRequest", string(data)}, " ")
}
