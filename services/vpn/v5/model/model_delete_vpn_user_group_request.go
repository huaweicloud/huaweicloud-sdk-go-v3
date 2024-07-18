package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpnUserGroupRequest Request Object
type DeleteVpnUserGroupRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 用户组ID
	GroupId string `json:"group_id"`
}

func (o DeleteVpnUserGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpnUserGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpnUserGroupRequest", string(data)}, " ")
}
