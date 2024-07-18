package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpnUserPasswordRequest Request Object
type UpdateVpnUserPasswordRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 用户ID
	UserId string `json:"user_id"`

	Body *UpdateVpnUserPasswordRequestBody `json:"body,omitempty"`
}

func (o UpdateVpnUserPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnUserPasswordRequest struct{}"
	}

	return strings.Join([]string{"UpdateVpnUserPasswordRequest", string(data)}, " ")
}
