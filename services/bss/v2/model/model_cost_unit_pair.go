package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CostUnitPair struct {

	// 成本单元名称。
	CostUnitName *string `json:"cost_unit_name,omitempty"`

	// 成本单元规则值。
	CostUnitRuleValue *string `json:"cost_unit_rule_value,omitempty"`
}

func (o CostUnitPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CostUnitPair struct{}"
	}

	return strings.Join([]string{"CostUnitPair", string(data)}, " ")
}
