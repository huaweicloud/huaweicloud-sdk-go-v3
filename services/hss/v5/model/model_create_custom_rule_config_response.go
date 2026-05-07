package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomRuleConfigResponse Response Object
type CreateCustomRuleConfigResponse struct {

	// **参数解释**： 规则ID **取值范围**： 字符长度1-36位
	RuleId         *string `json:"rule_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCustomRuleConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomRuleConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateCustomRuleConfigResponse", string(data)}, " ")
}
