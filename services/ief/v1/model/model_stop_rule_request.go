package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopRuleRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 规则ID
	RuleId string `json:"rule_id"`
}

func (o StopRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopRuleRequest struct{}"
	}

	return strings.Join([]string{"StopRuleRequest", string(data)}, " ")
}
