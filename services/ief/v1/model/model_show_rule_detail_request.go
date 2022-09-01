package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRuleDetailRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty" xml:"ief-instance-id"`

	// 规则ID
	RuleId string `json:"rule_id" xml:"rule_id"`
}

func (o ShowRuleDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRuleDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowRuleDetailRequest", string(data)}, " ")
}
