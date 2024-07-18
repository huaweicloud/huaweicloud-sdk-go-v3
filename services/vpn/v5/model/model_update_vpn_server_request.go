package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpnServerRequest Request Object
type UpdateVpnServerRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	Body *UpdateServerRequestBody `json:"body,omitempty"`
}

func (o UpdateVpnServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnServerRequest struct{}"
	}

	return strings.Join([]string{"UpdateVpnServerRequest", string(data)}, " ")
}
