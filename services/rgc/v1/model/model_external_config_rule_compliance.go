package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExternalConfigRuleCompliance struct {

	// 合规规则。
	RuleName *string `json:"rule_name,omitempty"`

	// 合规状态。
	Status *string `json:"status,omitempty"`

	// 外部规则ID。
	ControlId *string `json:"control_id,omitempty"`
}

func (o ExternalConfigRuleCompliance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalConfigRuleCompliance struct{}"
	}

	return strings.Join([]string{"ExternalConfigRuleCompliance", string(data)}, " ")
}
