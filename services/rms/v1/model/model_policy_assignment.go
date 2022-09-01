package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 规则
type PolicyAssignment struct {

	// 规则ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 规则名字
	Name *string `json:"name,omitempty" xml:"name"`

	// 规则描述
	Description *string `json:"description,omitempty" xml:"description"`

	PolicyFilter *PolicyFilterDefinition `json:"policy_filter,omitempty" xml:"policy_filter"`

	// 规则状态
	State *string `json:"state,omitempty" xml:"state"`

	// 规则创建时间
	Created *string `json:"created,omitempty" xml:"created"`

	// 规则更新时间
	Updated *string `json:"updated,omitempty" xml:"updated"`

	// 规则的策略ID
	PolicyDefinitionId *string `json:"policy_definition_id,omitempty" xml:"policy_definition_id"`

	// 规则参数
	Parameters map[string]PolicyParameterValue `json:"parameters,omitempty" xml:"parameters"`
}

func (o PolicyAssignment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyAssignment struct{}"
	}

	return strings.Join([]string{"PolicyAssignment", string(data)}, " ")
}
