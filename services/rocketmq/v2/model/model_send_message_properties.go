package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SendMessageProperties struct {

	// **参数解释**： 特性名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 特性值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o SendMessageProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendMessageProperties struct{}"
	}

	return strings.Join([]string{"SendMessageProperties", string(data)}, " ")
}
