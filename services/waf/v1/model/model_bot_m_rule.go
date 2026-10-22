package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BotMRule **参数解释：** BotM规则包含的信息，定义单条BotM规则的配置与状态。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type BotMRule struct {

	// **参数解释：** 规则ID，唯一标识当前BotM规则。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 规则名称，用于标识当前BotM规则的名称。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 规则描述，对当前BotM规则的功能说明。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 规则列表，当前BotM规则包含的具体检测特征。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Features *[]string `json:"features,omitempty"`

	// **参数解释：** 规则所属类别，标识规则的一级分类（如0表示基础检测类）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Category *int32 `json:"category,omitempty"`

	// **参数解释：** 规则所属子类别，标识规则的二级分类（如0表示已知Bot子类）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubCategory *int32 `json:"sub_category,omitempty"`

	// **参数解释：** 规则对应的防护动作，标识触发规则后执行的动作（如0表示放行）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DefenseAction *int32 `json:"defense_action,omitempty"`

	// **参数解释：** 规则创建的时间，时间戳（毫秒级）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CreatedTime *int64 `json:"created_time,omitempty"`

	// **参数解释：** 规则更新的时间，时间戳（毫秒级）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ModifiedTime *int64 `json:"modified_time,omitempty"`

	// **参数解释：** 规则目前是否被启用（true表示启用，false表示禁用）。 **约束限制：** 不涉及 **取值范围：** 仅支持true、false两个布尔值 **默认取值：** true
	Status *bool `json:"status,omitempty"`

	// **参数解释：** 交互置信度，标识主动交互检测的置信度阈值。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	InteractionConfidence *int32 `json:"interaction_confidence,omitempty"`
}

func (o BotMRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BotMRule struct{}"
	}

	return strings.Join([]string{"BotMRule", string(data)}, " ")
}
