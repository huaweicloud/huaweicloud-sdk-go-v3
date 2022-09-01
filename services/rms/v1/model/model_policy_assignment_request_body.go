package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 规则请求体
type PolicyAssignmentRequestBody struct {

	// 规则名字
	Name string `json:"name" xml:"name"`

	// 规则描述
	Description *string `json:"description,omitempty" xml:"description"`

	PolicyFilter *PolicyFilterDefinition `json:"policy_filter" xml:"policy_filter"`

	// 策略定义ID
	PolicyDefinitionId string `json:"policy_definition_id" xml:"policy_definition_id"`

	// 规则参数
	Parameters map[string]PolicyParameterValue `json:"parameters" xml:"parameters"`
}

func (o PolicyAssignmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyAssignmentRequestBody struct{}"
	}

	return strings.Join([]string{"PolicyAssignmentRequestBody", string(data)}, " ")
}
