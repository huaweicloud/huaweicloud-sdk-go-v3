package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRuleFixValuesInfo 用户键入的要修复的检查项的参数ID和参数值
type CheckRuleFixValuesInfo struct {

	// **参数解释**: 检查项的参数ID **约束限制**: 不涉及 **取值范围**: 字符串大小范围1-128 **默认取值**: 不涉及
	RuleParamId *int32 `json:"rule_param_id,omitempty"`

	// **参数解释**: 修复检查项参数具体值 **约束限制**: 不涉及 **取值范围**: 字符串大小范围0-512 **默认取值**: 不涉及
	FixValue *int32 `json:"fix_value,omitempty"`
}

func (o CheckRuleFixValuesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRuleFixValuesInfo struct{}"
	}

	return strings.Join([]string{"CheckRuleFixValuesInfo", string(data)}, " ")
}
