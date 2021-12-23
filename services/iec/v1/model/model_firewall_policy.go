package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 网络ACL策略。
type FirewallPolicy struct {
	// 网络ACL策略ID。

	Id string `json:"id"`
	// 网络ACL策略名称。

	Name *string `json:"name,omitempty"`
	// 网络ACL规则列表对象。

	FirewallRules []FirewallRule `json:"firewall_rules"`
	// ACL规则ID，表示在此ACL规则之后添加ACL规则

	InsertAfter *string `json:"insert_after,omitempty"`
	// ACL规则ID，表示在此ACL规则之前添加ACL规则

	InsertBefore *string `json:"insert_before,omitempty"`
}

func (o FirewallPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FirewallPolicy struct{}"
	}

	return strings.Join([]string{"FirewallPolicy", string(data)}, " ")
}
