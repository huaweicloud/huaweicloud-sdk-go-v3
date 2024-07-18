package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpnUserRequest Request Object
type DeleteVpnUserRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 用户ID
	UserId string `json:"user_id"`
}

func (o DeleteVpnUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpnUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpnUserRequest", string(data)}, " ")
}
