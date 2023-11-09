package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnConnectionRequest Request Object
type ShowVpnConnectionRequest struct {

	// vpn连接ID
	VpnConnectionId string `json:"vpn_connection_id"`
}

func (o ShowVpnConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowVpnConnectionRequest", string(data)}, " ")
}
