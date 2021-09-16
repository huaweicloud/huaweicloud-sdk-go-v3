package model

import (
	"encoding/json"

	"strings"
)

// 更新防火墙规则的参数
type UpdateFirewallRuleOption struct {
	EgressFirewallPolicy *FirewallPolicy `json:"egress_firewall_policy"`

	IngressFirewallPolicy *FirewallPolicy `json:"ingress_firewall_policy"`
}

func (o UpdateFirewallRuleOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateFirewallRuleOption struct{}"
	}

	return strings.Join([]string{"UpdateFirewallRuleOption", string(data)}, " ")
}
