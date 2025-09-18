package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginDtoOutputInfo struct {

	// **参数解释**： 插件输出配置的唯一标识。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 插件输出版本号。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 插件输出配置的描述。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 插件输出配置的源。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Source *string `json:"source,omitempty"`

	// **参数解释**： 插件输出配置的输出类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`
}

func (o PluginDtoOutputInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginDtoOutputInfo struct{}"
	}

	return strings.Join([]string{"PluginDtoOutputInfo", string(data)}, " ")
}
