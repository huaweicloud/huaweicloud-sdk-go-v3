package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnUserGroupRequest Request Object
type ShowVpnUserGroupRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 用户组ID
	GroupId string `json:"group_id"`
}

func (o ShowVpnUserGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnUserGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowVpnUserGroupRequest", string(data)}, " ")
}
