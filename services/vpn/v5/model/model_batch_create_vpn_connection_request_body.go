package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateVpnConnectionRequestBody struct {

	// VPN连接对象数组
	VpnConnections []CreateVpnConnectionRequestBodyContent `json:"vpn_connections"`
}

func (o BatchCreateVpnConnectionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateVpnConnectionRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateVpnConnectionRequestBody", string(data)}, " ")
}
