package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRuleRequest **参数解释：** 规则引擎配置内容 **约束限制：** 不涉及
type CreateRuleRequest struct {

	// **参数解释：** 规则名称 **约束限制：** 不涉及 **取值范围：** 1-50个字符 **默认取值：** 不涉及
	Name string `json:"name"`

	// **参数解释：** 是否开启规则 **约束限制：** 不涉及 **取值范围：** - on: 开启 - off: 关闭 **默认取值：** 不涉及
	Status string `json:"status"`

	// **参数解释：** 规则的优先级，数值越大，优先级越高 **约束限制：** 优先级不能重复 **取值范围：** 1-100 **默认取值：** 不涉及
	Priority int32 `json:"priority"`

	Conditions *Conditions `json:"conditions"`

	// **参数解释：** 满足规则条件后执行的动作 **约束限制：** 不涉及
	Actions []Actions `json:"actions"`
}

func (o CreateRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateRuleRequest", string(data)}, " ")
}
