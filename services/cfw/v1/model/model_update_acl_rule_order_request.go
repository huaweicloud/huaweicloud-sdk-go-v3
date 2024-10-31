package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAclRuleOrderRequest Request Object
type UpdateAclRuleOrderRequest struct {

	// 规则id，可通过[查询防护规则接口](ListAclRules.xml)查询获得，通过返回值中的data.records.rule_id（.表示各对象之间层级的区分）获得。
	AclRuleId string `json:"acl_rule_id"`

	// 企业项目ID，用户根据组织规划企业项目，对应的ID为企业项目ID，可通过[如何获取企业项目ID](cfw_02_0027.xml)获取，用户未开启企业项目时为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙id，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	Body *OrderRuleAclDto `json:"body,omitempty"`
}

func (o UpdateAclRuleOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclRuleOrderRequest struct{}"
	}

	return strings.Join([]string{"UpdateAclRuleOrderRequest", string(data)}, " ")
}
