package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateVpnConnectionResponse Response Object
type BatchCreateVpnConnectionResponse struct {

	// VPN连接对象数组
	VpnConnections *[]CreateResponseVpnConnection `json:"vpn_connections,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateVpnConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVpnConnectionResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateVpnConnectionResponse", string(data)}, " ")
}
