package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnConnectionsLogConfigRequest Request Object
type ShowVpnConnectionsLogConfigRequest struct {

	// P2C VPN网关实例ID
	P2cVgwId string `json:"p2c_vgw_id"`
}

func (o ShowVpnConnectionsLogConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnConnectionsLogConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowVpnConnectionsLogConfigRequest", string(data)}, " ")
}
