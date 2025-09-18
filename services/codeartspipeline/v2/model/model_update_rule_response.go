package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRuleResponse Response Object
type UpdateRuleResponse struct {

	// **参数解释**： 是否调用成功。 **取值范围**： - true：调用成功。 - false：调用失败。
	Status *bool `json:"status,omitempty"`

	// **参数解释**： 规则ID，规则的唯一标识，通过[分页获取规则列表](ListRule.xml)接口获取，data.id即为规则ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	RuleId         *string `json:"rule_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateRuleResponse", string(data)}, " ")
}
