package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetVpnConnectionRequest Request Object
type ResetVpnConnectionRequest struct {

	// vpn连接ID
	VpnConnectionId string `json:"vpn_connection_id"`
}

func (o ResetVpnConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetVpnConnectionRequest struct{}"
	}

	return strings.Join([]string{"ResetVpnConnectionRequest", string(data)}, " ")
}
