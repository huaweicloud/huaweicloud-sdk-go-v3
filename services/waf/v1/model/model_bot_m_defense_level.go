package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BotMDefenseLevel **参数解释：** BotM防护策略中的某一个防护等级，包含分数门限和对应防护动作 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type BotMDefenseLevel struct {

	// **参数解释：** 该防护等级对应的分数门限，触发当前防护等级的风险分数阈值 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Threshold *int32 `json:"threshold,omitempty"`

	// **参数解释：** 该防护等级对应的防护动作ID，如101表示拦截、102表示放行等 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DefenseAction *int32 `json:"defense_action,omitempty"`
}

func (o BotMDefenseLevel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BotMDefenseLevel struct{}"
	}

	return strings.Join([]string{"BotMDefenseLevel", string(data)}, " ")
}
