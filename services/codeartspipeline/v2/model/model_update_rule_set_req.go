package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRuleSetReq struct {

	// **参数解释**： 策略名称。 **约束限制**： 策略名称仅支持中文、大小写英文字母、数字、‘-’、‘_’。 **取值范围**： 不超过128个字符。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 策略中的规则列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Rules *[]UpdateRuleInstance `json:"rules,omitempty"`
}

func (o UpdateRuleSetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleSetReq struct{}"
	}

	return strings.Join([]string{"UpdateRuleSetReq", string(data)}, " ")
}
