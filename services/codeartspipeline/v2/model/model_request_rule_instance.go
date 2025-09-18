package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RequestRuleInstance struct {

	// **参数解释**： 规则实例ID **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 规则实例状态 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	IsValid *bool `json:"is_valid,omitempty"`
}

func (o RequestRuleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RequestRuleInstance struct{}"
	}

	return strings.Join([]string{"RequestRuleInstance", string(data)}, " ")
}
