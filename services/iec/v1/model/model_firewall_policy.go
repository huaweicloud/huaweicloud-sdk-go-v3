package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 网络ACL策略。
type FirewallPolicy struct {

	// 网络ACL策略ID。
	Id string `json:"id" xml:"id"`

	// 网络ACL策略名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 网络ACL规则列表对象。
	FirewallRules []FirewallRule `json:"firewall_rules" xml:"firewall_rules"`

	// ACL规则ID，表示在此ACL规则之后添加ACL规则
	InsertAfter *string `json:"insert_after,omitempty" xml:"insert_after"`

	// ACL规则ID，表示在此ACL规则之前添加ACL规则
	InsertBefore *string `json:"insert_before,omitempty" xml:"insert_before"`
}

func (o FirewallPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FirewallPolicy struct{}"
	}

	return strings.Join([]string{"FirewallPolicy", string(data)}, " ")
}
