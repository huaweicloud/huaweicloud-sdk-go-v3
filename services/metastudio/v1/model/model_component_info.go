package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComponentInfo 组件信息。
type ComponentInfo struct {

	// **参数解释**： 组件名称。 **约束限制**： 不涉及。 **取值范围**： 字符长度1-256位。 **默认取值**： 不涉及。
	ComponentName string `json:"component_name"`

	// **参数解释**： 组件类型。 **约束限制**： 不涉及。 **取值范围**： 字符长度1-256位。 **默认取值**： 不涉及。
	ComponentType string `json:"component_type"`

	// **参数解释**： 组件描述。 **约束限制**： 不涉及。 **取值范围**： 字符长度0-512位。 **默认取值**： 不涉及。
	ComponentDesc *string `json:"component_desc,omitempty"`
}

func (o ComponentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentInfo struct{}"
	}

	return strings.Join([]string{"ComponentInfo", string(data)}, " ")
}
