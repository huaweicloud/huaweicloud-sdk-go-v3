package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVirtualGatewayResponse Response Object
type CreateVirtualGatewayResponse struct {
	VirtualGateway *VirtualGateway `json:"virtual_gateway,omitempty"`

	// 操作请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVirtualGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirtualGatewayResponse struct{}"
	}

	return strings.Join([]string{"CreateVirtualGatewayResponse", string(data)}, " ")
}
