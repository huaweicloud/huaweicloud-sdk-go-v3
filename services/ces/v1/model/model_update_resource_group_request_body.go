package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateResourceGroupRequestBody 更新资源分组，请求参数。
type UpdateResourceGroupRequestBody struct {

	// **参数解释** 资源分组的名称 **约束限制** 不涉及 **取值范围** 只能为字母、数字、汉字、-或_，长度为[1,128]个字符 **默认取值** 不涉及
	GroupName string `json:"group_name"`

	// **参数解释** 手动创建时的资源详情。 **约束限制** 不超过1000个资源。
	Resources *[]CreateResourceGroup `json:"resources,omitempty"`

	// **参数解释** 资源分组添加资源方式 **约束限制** 不涉及 **取值范围** 取值只能为EPS（同步企业项目），TAG（标签动态匹配），不传为手动添加。 **默认取值** 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释** 该资源分组内包含的资源来源的企业项目ID。 **约束限制** type为EPS时必传，不超过50个企业项目ID。
	RelationIds *[]string `json:"relation_ids,omitempty"`

	// **参数解释** 标签动态匹配时的关联标签。 **约束限制** type为TAG时必传，不超过50个标签。
	Tags *[]ResourceGroupTagRelation `json:"tags,omitempty"`
}

func (o UpdateResourceGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResourceGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateResourceGroupRequestBody", string(data)}, " ")
}
