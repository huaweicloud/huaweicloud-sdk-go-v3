package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpnConnectionRequest Request Object
type UpdateVpnConnectionRequest struct {

	// vpn连接ID
	VpnConnectionId string `json:"vpn_connection_id"`

	Body *UpdateVpnConnectionRequestBody `json:"body,omitempty"`
}

func (o UpdateVpnConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnConnectionRequest struct{}"
	}

	return strings.Join([]string{"UpdateVpnConnectionRequest", string(data)}, " ")
}
