package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRuleInstance struct {

	// **参数解释**： 规则ID，规则的唯一标识，通过[分页获取规则列表](ListRule.xml)接口获取，data.id即为规则ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 规则启用状态。 **约束限制**： 不涉及。 **取值范围**： - true：启用。 - false：不启用。 **默认取值**： 不涉及。
	IsValid *bool `json:"is_valid,omitempty"`
}

func (o UpdateRuleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleInstance struct{}"
	}

	return strings.Join([]string{"UpdateRuleInstance", string(data)}, " ")
}
