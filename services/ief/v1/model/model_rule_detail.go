package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleDetail 规则配置
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
