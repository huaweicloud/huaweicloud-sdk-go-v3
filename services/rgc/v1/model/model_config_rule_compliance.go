package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigRuleCompliance struct {

	// 合规规则。
	RuleName *string `json:"rule_name,omitempty"`

	// 合规状态。
	Status *string `json:"status,omitempty"`

	// 区域信息。
	Region *string `json:"region,omitempty"`

	// 控制策略Id。
	ControlId *string `json:"control_id,omitempty"`
}

func (o ConfigRuleCompliance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigRuleCompliance struct{}"
	}

	return strings.Join([]string{"ConfigRuleCompliance", string(data)}, " ")
}
