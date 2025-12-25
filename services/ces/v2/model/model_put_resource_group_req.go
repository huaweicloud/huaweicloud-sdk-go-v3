package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutResourceGroupReq **参数解释** 资源分组修改请求体
type PutResourceGroupReq struct {

	// **参数解释** 资源分组名称。 **约束限制** 不涉及。 **取值范围** 包含字母、数字、_、-或汉字，长度为[1,128]个字符。 **默认取值** 不涉及。
	GroupName string `json:"group_name"`

	// **参数解释** 标签动态匹配时的关联标签。 **约束限制** type为TAG时必传，不超过50个标签。
	Tags *[]ResourceGroupTagRelation `json:"tags,omitempty"`

	// **参数解释** 资源匹配规则为组合匹配时传入的参数，已废弃。 **约束限制** type为COMB时传入，不超过50个条件。
	EnterpriseProjectIdAndTags *[]EnterpriseProjectIdAndTags `json:"enterprise_project_id_and_tags,omitempty"`

	// **参数解释** 匹配企业项目时关联的企业项目ID。 **约束限制** type为EPS时必传，不超过50个企业项目。
	ExtendRelationIds *[]string `json:"extend_relation_ids,omitempty"`

	// **参数解释** 实例名称匹配参数。 **约束限制** type为NAME时必传，不超过50个实例。
	Instances *[]Instance `json:"instances,omitempty"`

	// **参数解释** 资源层级为云产品时的云产品名称 **约束限制** 不涉及。 **取值范围** 一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"。多个云产品则用“;”隔开，如\"SERVICE.BMS,instance_id;SYS.ECS,instance_id\"。不超过10240个字符。 **默认取值** 不涉及。
	ProductNames *string `json:"product_names,omitempty"`

	CombRelation *CombRelation `json:"comb_relation,omitempty"`
}

func (o PutResourceGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutResourceGroupReq struct{}"
	}

	return strings.Join([]string{"PutResourceGroupReq", string(data)}, " ")
}
