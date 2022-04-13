package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePolicyAssignmentsResponse struct {
	// 规则ID

	Id *string `json:"id,omitempty"`
	// 规则名字

	Name *string `json:"name,omitempty"`
	// 规则描述

	Description *string `json:"description,omitempty"`

	PolicyFilter *PolicyFilterDefinition `json:"policy_filter,omitempty"`
	// 规则状态

	State *string `json:"state,omitempty"`
	// 规则创建时间

	Created *string `json:"created,omitempty"`
	// 规则更新时间

	Updated *string `json:"updated,omitempty"`
	// 规则的策略ID

	PolicyDefinitionId *string `json:"policy_definition_id,omitempty"`
	// 规则参数

	Parameters     map[string]PolicyParameterValue `json:"parameters,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o CreatePolicyAssignmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePolicyAssignmentsResponse struct{}"
	}

	return strings.Join([]string{"CreatePolicyAssignmentsResponse", string(data)}, " ")
}
