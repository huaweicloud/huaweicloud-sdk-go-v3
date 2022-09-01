package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type NeutronRemoveFirewallRuleResponse struct {

	// 功能说明：网络ACL策略ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 功能说明：网络ACL策略名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 功能说明：网络ACL策略的描述信息 取值范围：0-255个字符
	Description *string `json:"description,omitempty" xml:"description"`

	// 功能说明：网络ACL策略关联的规则的ID列表
	FirewallRules *[]string `json:"firewall_rules,omitempty" xml:"firewall_rules"`

	// 功能说明：每次policy或者它相关的rule有变动，该参数将会被置为False
	Audited *bool `json:"audited,omitempty" xml:"audited"`

	// 功能说明：网络ACL策略是否对其他网络ACL策略可见 取值范围：true(可见)、false(不可见)
	Public *bool `json:"public,omitempty" xml:"public"`

	// 功能说明：网络ACL策略所属项目ID
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id"`

	// 功能说明：网络ACL策略所属项目ID
	ProjectId      *string `json:"project_id,omitempty" xml:"project_id"`
	HttpStatusCode int     `json:"-"`
}

func (o NeutronRemoveFirewallRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronRemoveFirewallRuleResponse struct{}"
	}

	return strings.Join([]string{"NeutronRemoveFirewallRuleResponse", string(data)}, " ")
}
