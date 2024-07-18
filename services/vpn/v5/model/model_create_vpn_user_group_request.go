package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpnUserGroupRequest Request Object
type CreateVpnUserGroupRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 幂等性标识
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *CreateVpnUserGroupRequestBody `json:"body,omitempty"`
}

func (o CreateVpnUserGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnUserGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateVpnUserGroupRequest", string(data)}, " ")
}
