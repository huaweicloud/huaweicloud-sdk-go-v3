package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Rule
type Rule struct {

	// 候选集表名。
	TableName string `json:"table_name"`

	// 规则占比。
	RuleRatio int32 `json:"rule_ratio"`

	// 优先级。
	Priority int32 `json:"priority"`
}

func (o Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Rule struct{}"
	}

	return strings.Join([]string{"Rule", string(data)}, " ")
}
