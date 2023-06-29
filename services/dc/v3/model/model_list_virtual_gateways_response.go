package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVirtualGatewaysResponse Response Object
type ListVirtualGatewaysResponse struct {

	// 操作请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 虚拟网关对象列表
	VirtualGateways *[]VirtualGateway `json:"virtual_gateways,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListVirtualGatewaysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVirtualGatewaysResponse struct{}"
	}

	return strings.Join([]string{"ListVirtualGatewaysResponse", string(data)}, " ")
}
