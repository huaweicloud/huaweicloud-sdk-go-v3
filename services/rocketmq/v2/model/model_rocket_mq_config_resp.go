package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RocketMqConfigResp struct {

	// **参数解释**： RocketMQ配置名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： RocketMQ配置当前值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释**： RocketMQ配置的类型。 **约束限制**： 不涉及。 **取值范围**： - dynamic：动态。 - static：静态。    **默认取值**： 不涉及。
	ConfigType *string `json:"config_type,omitempty"`

	// **参数解释**： RocketMQ配置的默认值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DefaultValue *string `json:"default_value,omitempty"`

	// **参数解释**： RocketMQ配置取值的范围。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ValidValues *string `json:"valid_values,omitempty"`

	// **参数解释**： RocketMQ配置值的类型。 **约束限制**： 不涉及。 **取值范围**： - integer：整数类型。 - boolean：布尔类型。 **默认取值**： 不涉及。
	ValueType *string `json:"value_type,omitempty"`
}

func (o RocketMqConfigResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RocketMqConfigResp struct{}"
	}

	return strings.Join([]string{"RocketMqConfigResp", string(data)}, " ")
}
