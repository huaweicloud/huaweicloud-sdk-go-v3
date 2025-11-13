package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveFirewallRulesRequestBody This is a auto create Body Object
type RemoveFirewallRulesRequestBody struct {
	Firewall *FirewallRemoveRuleOption `json:"firewall"`

	// 功能说明：是否只预检此次请求 取值范围： -true：发送检查请求，不会删除网络ACL规则。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回响应码202。 -false（默认值）：发送正常请求，并直接删除网络ACL规则。
	DryRun *bool `json:"dry_run,omitempty"`
}

func (o RemoveFirewallRulesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveFirewallRulesRequestBody struct{}"
	}

	return strings.Join([]string{"RemoveFirewallRulesRequestBody", string(data)}, " ")
}
