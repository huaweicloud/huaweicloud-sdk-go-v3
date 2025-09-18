package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StagePluginsQueryDto struct {

	// **参数解释**： 用于区分插件为流水线可使用/模板可使用。 **约束限制**： 不涉及。 **取值范围**： - pipeline：流水线可使用。 - template：模板可使用。 **默认取值**： 不涉及。
	UseCondition *string `json:"use_condition,omitempty"`

	// **参数解释**： 微服务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CompId *string `json:"comp_id,omitempty"`

	// **参数解释**： 微服务名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CompName *string `json:"comp_name,omitempty"`

	// **参数解释**： 局点ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CloudId *string `json:"cloud_id,omitempty"`

	// **参数解释**： 策略ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StrategyId *string `json:"strategy_id,omitempty"`

	// **参数解释**： 流水线类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Category *string `json:"category,omitempty"`

	// **参数解释**： 是否发布流水线。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PublishTab *string `json:"publish_tab,omitempty"`

	// **参数解释**： 部署平台。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Platform *string `json:"platform,omitempty"`

	// **参数解释**： 组件类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CompExtendType *string `json:"comp_extend_type,omitempty"`

	// **参数解释**： 部署类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DeployType *string `json:"deploy_type,omitempty"`
}

func (o StagePluginsQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StagePluginsQueryDto struct{}"
	}

	return strings.Join([]string{"StagePluginsQueryDto", string(data)}, " ")
}
