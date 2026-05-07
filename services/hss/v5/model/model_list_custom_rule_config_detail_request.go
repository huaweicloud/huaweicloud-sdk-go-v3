package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomRuleConfigDetailRequest Request Object
type ListCustomRuleConfigDetailRequest struct {

	// **参数解释**： 规则ID **约束限制**： 必填 **取值范围**： 字符长度1-36位 **默认取值**： 不涉及
	RuleId string `json:"rule_id"`
}

func (o ListCustomRuleConfigDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomRuleConfigDetailRequest struct{}"
	}

	return strings.Join([]string{"ListCustomRuleConfigDetailRequest", string(data)}, " ")
}
