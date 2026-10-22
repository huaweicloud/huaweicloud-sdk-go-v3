package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBotMRulesResponse Response Object
type ListBotMRulesResponse struct {

	// **参数解释：** 策略Id，关联BotM规则的防护策略唯一标识。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释：** 租户Id，当前BotM规则所属的租户唯一标识。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释：** 已知Bot相关的所有规则，包含针对已知Bot的检测与防护规则。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	KnownBotDetection *[]BotMRule `json:"known_bot_detection,omitempty"`

	// **参数解释：** 透明检测相关的所有规则，包含无感知的Bot透明检测规则。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TransparentDetection *[]BotMRule `json:"transparent_detection,omitempty"`

	BehaviorDetection *BotMBehaviorDetectionRule `json:"behavior_detection,omitempty"`

	// **参数解释：** 需要BOT检测的流量检测条件，定义触发Bot检测的流量筛选规则。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	TrafficDetectionConditions *[]TrafficDetectionConditionDto `json:"traffic_detection_conditions,omitempty"`

	// **参数解释：** 主动特征检测规则列表，包含需要主动交互验证的Bot检测规则。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	InteractiveDetection *[]BotMRule `json:"interactive_detection,omitempty"`
	HttpStatusCode       int         `json:"-"`
}

func (o ListBotMRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBotMRulesResponse struct{}"
	}

	return strings.Join([]string{"ListBotMRulesResponse", string(data)}, " ")
}
