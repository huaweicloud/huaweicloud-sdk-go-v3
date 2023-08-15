package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVirtualGatewayResponse Response Object
type ShowVirtualGatewayResponse struct {
	VirtualGateway *VirtualGateway `json:"virtual_gateway,omitempty"`

	// 操作请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowVirtualGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVirtualGatewayResponse struct{}"
	}

	return strings.Join([]string{"ShowVirtualGatewayResponse", string(data)}, " ")
}
