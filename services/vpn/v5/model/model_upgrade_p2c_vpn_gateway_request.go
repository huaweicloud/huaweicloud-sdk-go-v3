package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeP2cVpnGatewayRequest Request Object
type UpgradeP2cVpnGatewayRequest struct {

	// P2C VPN网关实例ID
	P2cVgwId string `json:"p2c_vgw_id"`

	Body *UpgradeRequestBody `json:"body,omitempty"`
}

func (o UpgradeP2cVpnGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeP2cVpnGatewayRequest struct{}"
	}

	return strings.Join([]string{"UpgradeP2cVpnGatewayRequest", string(data)}, " ")
}
