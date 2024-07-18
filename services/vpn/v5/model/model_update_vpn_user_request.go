package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpnUserRequest Request Object
type UpdateVpnUserRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 用户ID
	UserId string `json:"user_id"`

	Body *UpdateVpnUserRequestBody `json:"body,omitempty"`
}

func (o UpdateVpnUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnUserRequest struct{}"
	}

	return strings.Join([]string{"UpdateVpnUserRequest", string(data)}, " ")
}
