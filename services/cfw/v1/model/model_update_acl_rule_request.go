package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAclRuleRequest Request Object
type UpdateAclRuleRequest struct {

	// **参数解释**： 规则ID，可通过[查询防护规则接口](ListAclRules.xml)查询获得，通过返回值中的data.records.rule_id（.表示各对象之间层级的区分）获得。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AclRuleId string `json:"acl_rule_id"`

	// **参数解释**： 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 防火墙ID，用户创建防火墙实例后产生的唯一ID，配置后可区分不同防火墙，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取 **约束限制**： 不涉及 **取值范围**： 32位UUID **默认取值**： 不涉及
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	Body *UpdateRuleAclDto `json:"body,omitempty"`
}

func (o UpdateAclRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateAclRuleRequest", string(data)}, " ")
}
