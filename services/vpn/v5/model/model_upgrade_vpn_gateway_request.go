package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeVpnGatewayRequest Request Object
type UpgradeVpnGatewayRequest struct {

	// VPN网关实例ID
	VgwId string `json:"vgw_id"`

	Body *UpgradeRequestBody `json:"body,omitempty"`
}

func (o UpgradeVpnGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeVpnGatewayRequest struct{}"
	}

	return strings.Join([]string{"UpgradeVpnGatewayRequest", string(data)}, " ")
}
