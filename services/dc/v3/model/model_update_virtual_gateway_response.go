package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateVirtualGatewayResponse struct {
	VirtualGateway *VirtualGateway `json:"virtual_gateway,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o UpdateVirtualGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualGatewayResponse struct{}"
	}

	return strings.Join([]string{"UpdateVirtualGatewayResponse", string(data)}, " ")
}
