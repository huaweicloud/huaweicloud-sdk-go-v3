package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBuiltInPolicyDefinitionResponse struct {

	// 策略id
	Id *string `json:"id,omitempty" xml:"id"`

	// 策略名字
	Name *string `json:"name,omitempty" xml:"name"`

	// 策略类型
	PolicyType *string `json:"policy_type,omitempty" xml:"policy_type"`

	// 策略描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 策略语法类型
	PolicyRuleType *string `json:"policy_rule_type,omitempty" xml:"policy_rule_type"`

	// 策略规则
	PolicyRule *interface{} `json:"policy_rule,omitempty" xml:"policy_rule"`

	// 关键词列表
	Keywords *[]string `json:"keywords,omitempty" xml:"keywords"`

	// 策略参数
	Parameters     map[string]PolicyParameterDefinition `json:"parameters,omitempty" xml:"parameters"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ShowBuiltInPolicyDefinitionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuiltInPolicyDefinitionResponse struct{}"
	}

	return strings.Join([]string{"ShowBuiltInPolicyDefinitionResponse", string(data)}, " ")
}
