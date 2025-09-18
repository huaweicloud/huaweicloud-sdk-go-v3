package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRuleResponse Response Object
type CreateRuleResponse struct {

	// **参数解释**： 创建状态。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Status *bool `json:"status,omitempty"`

	// **参数解释**： 规则ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RuleId         *string `json:"rule_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateRuleResponse", string(data)}, " ")
}
