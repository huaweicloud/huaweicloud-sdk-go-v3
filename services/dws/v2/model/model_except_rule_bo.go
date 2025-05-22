package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExceptRuleBo **参数解释**： 异常规则信息对象。 **取值范围**： 不涉及。
type ExceptRuleBo struct {

	// **参数解释**： 规则名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 触发异常规则操作。 **取值范围**： 不涉及。
	Action *string `json:"action,omitempty"`

	// **参数解释**： 异常规则绑定的资源池名称列表。 **取值范围**： 不涉及。
	Queues *[]string `json:"queues,omitempty"`

	// **参数解释**： 异常规则配置项。 **取值范围**： 不涉及。
	ExceptRules map[string]string `json:"except_rules,omitempty"`
}

func (o ExceptRuleBo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExceptRuleBo struct{}"
	}

	return strings.Join([]string{"ExceptRuleBo", string(data)}, " ")
}
