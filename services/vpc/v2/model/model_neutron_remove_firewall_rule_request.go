package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronRemoveFirewallRuleRequest Request Object
type NeutronRemoveFirewallRuleRequest struct {

	// 网络ACL防火墙策略ID
	FirewallPolicyId string `json:"firewall_policy_id"`

	Body *NeutronRemoveFirewallRuleRequestBody `json:"body,omitempty"`
}

func (o NeutronRemoveFirewallRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronRemoveFirewallRuleRequest struct{}"
	}

	return strings.Join([]string{"NeutronRemoveFirewallRuleRequest", string(data)}, " ")
}
