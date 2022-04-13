package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 在IoTDA实例中的规则ID及其配套的推送数据动作ID
type IotdaRuleAction struct {
	// IoTDA中的规则Id

	RuleId string `json:"rule_id"`
	// IoTDA中推送数据动作ID

	ActionId string `json:"action_id"`
}

func (o IotdaRuleAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IotdaRuleAction struct{}"
	}

	return strings.Join([]string{"IotdaRuleAction", string(data)}, " ")
}
