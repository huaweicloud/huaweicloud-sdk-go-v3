package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetVpnUserPasswordRequest Request Object
type ResetVpnUserPasswordRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 用户ID
	UserId string `json:"user_id"`

	Body *ResetVpnUserPasswordRequestBody `json:"body,omitempty"`
}

func (o ResetVpnUserPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetVpnUserPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetVpnUserPasswordRequest", string(data)}, " ")
}
