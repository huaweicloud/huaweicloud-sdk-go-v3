package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRuleRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 规则ID
	RuleId string `json:"rule_id" xml:"rule_id"`
}

func (o ShowRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowRuleRequest", string(data)}, " ")
}
