package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowActionRuleRequest Request Object
type ShowActionRuleRequest struct {

	// 告警规则名称
	RuleName string `json:"rule_name"`
}

func (o ShowActionRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowActionRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowActionRuleRequest", string(data)}, " ")
}
