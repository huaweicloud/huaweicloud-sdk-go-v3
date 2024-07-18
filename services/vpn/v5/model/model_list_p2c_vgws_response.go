package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListP2cVgwsResponse Response Object
type ListP2cVgwsResponse struct {

	// 网关信息
	P2cVpnGateways *[]ShowResponseP2cVgw `json:"p2c_vpn_gateways,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListP2cVgwsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListP2cVgwsResponse struct{}"
	}

	return strings.Join([]string{"ListP2cVgwsResponse", string(data)}, " ")
}
