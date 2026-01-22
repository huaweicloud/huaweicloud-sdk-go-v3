package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RocketMqProductSupportFeaturesEntity **参数解释**： 支持的特性功能。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type RocketMqProductSupportFeaturesEntity struct {

	// **参数解释**： 特性名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 功能特性的键值对。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Properties map[string]string `json:"properties,omitempty"`
}

func (o RocketMqProductSupportFeaturesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RocketMqProductSupportFeaturesEntity struct{}"
	}

	return strings.Join([]string{"RocketMqProductSupportFeaturesEntity", string(data)}, " ")
}
