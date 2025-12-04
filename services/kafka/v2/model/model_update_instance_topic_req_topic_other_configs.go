package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceTopicReqTopicOtherConfigs struct {

	// **参数解释**： 配置名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 配置值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o UpdateInstanceTopicReqTopicOtherConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceTopicReqTopicOtherConfigs struct{}"
	}

	return strings.Join([]string{"UpdateInstanceTopicReqTopicOtherConfigs", string(data)}, " ")
}
