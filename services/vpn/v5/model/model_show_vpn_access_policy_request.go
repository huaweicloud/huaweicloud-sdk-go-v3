package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnAccessPolicyRequest Request Object
type ShowVpnAccessPolicyRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// VPN访问策略ID
	PolicyId string `json:"policy_id"`
}

func (o ShowVpnAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowVpnAccessPolicyRequest", string(data)}, " ")
}
