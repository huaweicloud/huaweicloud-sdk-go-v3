package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVgwsResponse Response Object
type ListVgwsResponse struct {

	// 网关信息
	VpnGateways *[]ResponseVpnGateway `json:"vpn_gateways,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListVgwsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVgwsResponse struct{}"
	}

	return strings.Join([]string{"ListVgwsResponse", string(data)}, " ")
}
