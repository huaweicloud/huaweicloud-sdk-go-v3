package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutResourceGroupReq 资源分组修改请求体
type PutResourceGroupReq struct {

	// 资源分组名称，只能为字母、数字、汉字、-、_，最大长度为128
	GroupName string `json:"group_name"`

	// 标签动态匹配时的关联标签,type为TAG时该字段不为空
	Tags *[]ResourceGroupTagRelation `json:"tags,omitempty"`

	// 资源匹配规则为组合匹配时传入的参数
	EnterpriseProjectIdAndTags *[]EnterpriseProjectIdAndTags `json:"enterprise_project_id_and_tags,omitempty"`

	// 智能添加时企业项目匹配传入参数
	ExtendRelationIds *[]string `json:"extend_relation_ids,omitempty"`

	// 实例名称匹配参数
	Instances *[]Instance `json:"instances,omitempty"`

	// 修改资源层级为云产品时的云产品的取值，一般由\"服务命名空间,服务首层维度名称\"组成，如\"SYS.ECS,instance_id\"。多个云产品则用“;”隔开，如\"SERVICE.BMS,instance_id;SYS.ECS,instance_id\"。
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
