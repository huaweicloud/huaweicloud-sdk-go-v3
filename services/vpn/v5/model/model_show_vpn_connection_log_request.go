package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnConnectionLogRequest Request Object
type ShowVpnConnectionLogRequest struct {

	// vpn连接ID
	VpnConnectionId string `json:"vpn_connection_id"`
}

func (o ShowVpnConnectionLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnConnectionLogRequest struct{}"
	}

	return strings.Join([]string{"ShowVpnConnectionLogRequest", string(data)}, " ")
}
