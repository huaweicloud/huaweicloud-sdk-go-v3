package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type NeutronShowFirewallPolicyRequest struct {
	// 网络ACL防火墙策略ID

	FirewallPolicyId string `json:"firewall_policy_id"`
}

func (o NeutronShowFirewallPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowFirewallPolicyRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowFirewallPolicyRequest", string(data)}, " ")
}
