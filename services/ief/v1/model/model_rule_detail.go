package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 规则配置
type RuleDetail struct {
	Rule *RuleConfig `json:"rule,omitempty"`
}

func (o RuleDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleDetail struct{}"
	}

	return strings.Join([]string{"RuleDetail", string(data)}, " ")
}
