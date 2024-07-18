package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddVpnUsersToGroupRequest Request Object
type AddVpnUsersToGroupRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 用户组ID
	GroupId string `json:"group_id"`

	Body *AddVpnUserToGroupRequestBody `json:"body,omitempty"`
}

func (o AddVpnUsersToGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVpnUsersToGroupRequest struct{}"
	}

	return strings.Join([]string{"AddVpnUsersToGroupRequest", string(data)}, " ")
}
