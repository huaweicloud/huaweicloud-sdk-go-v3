package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpnUserRequest Request Object
type CreateVpnUserRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 幂等性标识
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *CreateVpnUserRequestBody `json:"body,omitempty"`
}

func (o CreateVpnUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnUserRequest struct{}"
	}

	return strings.Join([]string{"CreateVpnUserRequest", string(data)}, " ")
}
