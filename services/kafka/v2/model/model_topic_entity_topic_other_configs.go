package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopicEntityTopicOtherConfigs struct {

	// **参数解释**： 配置名称。 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 配置有效值。 **取值范围**： 不涉及
	ValidValues *string `json:"valid_values,omitempty"`

	// **参数解释**： 配置默认值。 **取值范围**： 不涉及
	DefaultValue *string `json:"default_value,omitempty"`

	// **参数解释**： 配置类型。 **取值范围**： - dynamic：动态。 - static：静态。
	ConfigType *string `json:"config_type,omitempty"`

	// **参数解释**： 配置值。 **取值范围**： 不涉及
	Value *string `json:"value,omitempty"`

	// **参数解释**： 配置值类型。 **取值范围**： 不涉及
	ValueType *string `json:"value_type,omitempty"`
}

func (o TopicEntityTopicOtherConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicEntityTopicOtherConfigs struct{}"
	}

	return strings.Join([]string{"TopicEntityTopicOtherConfigs", string(data)}, " ")
}
