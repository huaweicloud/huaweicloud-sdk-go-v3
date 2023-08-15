package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVirtualGatewayRequestBody struct {
	VirtualGateway *UpdateVirtualGateway `json:"virtual_gateway,omitempty"`
}

func (o UpdateVirtualGatewayRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualGatewayRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVirtualGatewayRequestBody", string(data)}, " ")
}
