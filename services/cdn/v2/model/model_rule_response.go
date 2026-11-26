package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleResponse 查询规则请求体
type RuleResponse struct {

	// **参数解释：** 规则id **约束限制：** 不涉及 **取值范围：** 不涉及
	RuleId string `json:"rule_id"`

	// **参数解释：** 规则名称 **约束限制：** 不涉及 **取值范围：** 不涉及
	Name string `json:"name"`

	// 规则状态，on：开启，off：关闭。
	Status string `json:"status"`

	// **参数解释：** 此条规则的优先级，数值越大，优先级越高。 **约束限制：** 不涉及 **取值范围：** 1-1000
	Priority int32 `json:"priority"`

	Conditions *Conditions `json:"conditions"`

	// 满足规则条件后执行的动作列表
	Actions []Actions `json:"actions"`
}

func (o RuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleResponse struct{}"
	}

	return strings.Join([]string{"RuleResponse", string(data)}, " ")
}
