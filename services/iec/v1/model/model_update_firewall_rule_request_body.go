package model

import (
	"encoding/json"

	"strings"
)

// 更新防火墙规则请求体
type UpdateFirewallRuleRequestBody struct {
	Firewall *UpdateFirewallRuleOption `json:"firewall,omitempty"`
}

func (o UpdateFirewallRuleRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateFirewallRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFirewallRuleRequestBody", string(data)}, " ")
}
