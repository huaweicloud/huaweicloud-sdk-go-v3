package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateVirtualGatewayRequest struct {

	// 虚拟网关ID
	VirtualGatewayId string `json:"virtual_gateway_id"`

	Body *UpdateVirtualGatewayRequestBody `json:"body,omitempty"`
}

func (o UpdateVirtualGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualGatewayRequest struct{}"
	}

	return strings.Join([]string{"UpdateVirtualGatewayRequest", string(data)}, " ")
}
