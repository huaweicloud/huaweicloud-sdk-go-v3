package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpnUserGroupRequest Request Object
type UpdateVpnUserGroupRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 用户组ID
	GroupId string `json:"group_id"`

	Body *UpdateVpnUserGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateVpnUserGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnUserGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateVpnUserGroupRequest", string(data)}, " ")
}
