package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpnConnectionsLogConfigRequest Request Object
type UpdateVpnConnectionsLogConfigRequest struct {

	// P2C VPN网关实例ID
	P2cVgwId string `json:"p2c_vgw_id"`

	Body *UpdateVpnLogConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateVpnConnectionsLogConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnConnectionsLogConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateVpnConnectionsLogConfigRequest", string(data)}, " ")
}
