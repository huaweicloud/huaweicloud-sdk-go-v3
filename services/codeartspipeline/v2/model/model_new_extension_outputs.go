package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NewExtensionOutputs struct {

	// **参数解释**： 名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 扩展信息定义。 **取值范围**： 不涉及。
	Prop map[string]string `json:"prop,omitempty"`
}

func (o NewExtensionOutputs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NewExtensionOutputs struct{}"
	}

	return strings.Join([]string{"NewExtensionOutputs", string(data)}, " ")
}
