package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFirewallRuleRequestBody 更新防火墙规则请求体
type UpdateFirewallRuleRequestBody struct {
	Firewall *UpdateFirewallRuleOption `json:"firewall,omitempty"`
}

func (o UpdateFirewallRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFirewallRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFirewallRuleRequestBody", string(data)}, " ")
}
