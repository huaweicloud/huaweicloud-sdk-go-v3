package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpnConnectionsLogConfigRequest Request Object
type DeleteVpnConnectionsLogConfigRequest struct {

	// P2C VPN网关实例ID
	P2cVgwId string `json:"p2c_vgw_id"`
}

func (o DeleteVpnConnectionsLogConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpnConnectionsLogConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpnConnectionsLogConfigRequest", string(data)}, " ")
}
