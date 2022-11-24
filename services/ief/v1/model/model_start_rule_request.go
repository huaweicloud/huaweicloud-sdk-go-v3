package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartRuleRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 规则ID
	RuleId string `json:"rule_id"`
}

func (o StartRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartRuleRequest struct{}"
	}

	return strings.Join([]string{"StartRuleRequest", string(data)}, " ")
}
