package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Rule 规则信息。
type Rule struct {
	Scope *RuleScope `json:"scope"`

	ProductRule *ProductRule `json:"product_rule,omitempty"`

	PathRule *PathRule `json:"path_rule,omitempty"`
}

func (o Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Rule struct{}"
	}

	return strings.Join([]string{"Rule", string(data)}, " ")
}
