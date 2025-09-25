package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BotMBehaviorDetectionRule **参数解释：** BotM行为检测规则包含的信息，关联行为检测的基础规则与防护策略。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type BotMBehaviorDetectionRule struct {
	Rule *BotMRule `json:"rule,omitempty"`

	DefenseStrategy *BotMDefenseStrategy `json:"defense_strategy,omitempty"`
}

func (o BotMBehaviorDetectionRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BotMBehaviorDetectionRule struct{}"
	}

	return strings.Join([]string{"BotMBehaviorDetectionRule", string(data)}, " ")
}
