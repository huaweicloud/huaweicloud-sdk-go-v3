package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveVpnUsersFromGroupRequest Request Object
type RemoveVpnUsersFromGroupRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 用户组ID
	GroupId string `json:"group_id"`

	Body *RemoveVpnUserFromGroupRequestBody `json:"body,omitempty"`
}

func (o RemoveVpnUsersFromGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveVpnUsersFromGroupRequest struct{}"
	}

	return strings.Join([]string{"RemoveVpnUsersFromGroupRequest", string(data)}, " ")
}
