package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnUserRequest Request Object
type ShowVpnUserRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 用户ID
	UserId string `json:"user_id"`
}

func (o ShowVpnUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnUserRequest struct{}"
	}

	return strings.Join([]string{"ShowVpnUserRequest", string(data)}, " ")
}
