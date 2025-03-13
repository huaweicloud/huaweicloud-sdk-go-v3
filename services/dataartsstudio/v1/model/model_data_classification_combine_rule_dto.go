package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataClassificationCombineRuleDto struct {

	// 规则名称
	Name string `json:"name"`

	// 密级ID
	SecrecyLevelId string `json:"secrecy_level_id"`

	// 条件表达式
	CombineExpression string `json:"combine_expression"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 分类ID
	CategoryId *string `json:"category_id,omitempty"`

	// 使能状态。
	Enable *bool `json:"enable,omitempty"`

	// 条件单规则
	SingleExpressions []DataClassificationSingleRuleDto `json:"single_expressions"`
}

func (o DataClassificationCombineRuleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassificationCombineRuleDto struct{}"
	}

	return strings.Join([]string{"DataClassificationCombineRuleDto", string(data)}, " ")
}
