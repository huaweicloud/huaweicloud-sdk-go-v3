package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpnAccessPolicyRequest Request Object
type DeleteVpnAccessPolicyRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// VPN访问策略ID
	PolicyId string `json:"policy_id"`
}

func (o DeleteVpnAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpnAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteVpnAccessPolicyRequest", string(data)}, " ")
}
