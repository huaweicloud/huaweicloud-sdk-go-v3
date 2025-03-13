package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataClassificationCombineRuleCheckDto struct {

	// 条件表达式
	Expression *string `json:"expression,omitempty"`

	CombineInputData *DataClassificationCombineRuleCheckDtoCombineInputData `json:"combine_input_data,omitempty"`

	// 分类ID
	Combine *bool `json:"combine,omitempty"`

	// 条件单规则列表
	SingleRuleCheckList *[]DataClassificationSingleRuleDto `json:"single_rule_check_list,omitempty"`
}

func (o DataClassificationCombineRuleCheckDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassificationCombineRuleCheckDto struct{}"
	}

	return strings.Join([]string{"DataClassificationCombineRuleCheckDto", string(data)}, " ")
}
