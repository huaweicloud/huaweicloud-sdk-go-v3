package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVirtualGatewayResponse Response Object
type UpdateVirtualGatewayResponse struct {
	VirtualGateway *VirtualGateway `json:"virtual_gateway,omitempty"`

	// 操作请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateVirtualGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualGatewayResponse struct{}"
	}

	return strings.Join([]string{"UpdateVirtualGatewayResponse", string(data)}, " ")
}
