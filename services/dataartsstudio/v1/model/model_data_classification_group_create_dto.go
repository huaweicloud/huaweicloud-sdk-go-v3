package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataClassificationGroupCreateDto struct {

	// 规则名称
	Name string `json:"name"`

	// 规则id列表
	RuleIds []string `json:"rule_ids"`

	// 规则组描述
	Description *string `json:"description,omitempty"`

	// 需要创建的规则
	CreateRules *[]DataClassificationGroupCombineRuleDto `json:"create_rules,omitempty"`
}

func (o DataClassificationGroupCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassificationGroupCreateDto struct{}"
	}

	return strings.Join([]string{"DataClassificationGroupCreateDto", string(data)}, " ")
}
