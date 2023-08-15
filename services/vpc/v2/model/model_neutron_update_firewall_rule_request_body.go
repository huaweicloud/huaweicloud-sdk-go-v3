package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateFirewallRuleRequestBody
type NeutronUpdateFirewallRuleRequestBody struct {
	FirewallRule *NeutronUpdateFirewallRuleOption `json:"firewall_rule"`
}

func (o NeutronUpdateFirewallRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallRuleRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallRuleRequestBody", string(data)}, " ")
}
