package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpnAccessPolicyRequest Request Object
type UpdateVpnAccessPolicyRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// VPN访问策略ID
	PolicyId string `json:"policy_id"`

	Body *UpdateVpnAccessPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateVpnAccessPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnAccessPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateVpnAccessPolicyRequest", string(data)}, " ")
}
