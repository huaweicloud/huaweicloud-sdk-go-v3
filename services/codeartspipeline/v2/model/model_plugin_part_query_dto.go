package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginPartQueryDto struct {

	// **参数解释**： 插件名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PluginName string `json:"plugin_name"`

	// **参数解释**： 展示名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DisplayName string `json:"display_name"`

	// **参数解释**： 版本号。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Version string `json:"version"`

	// **参数解释**： 插件属性。 **约束限制**： 不涉及。 **取值范围**： - custom：自定义插件。 - official：官方插件。 **默认取值**： 不涉及。
	PluginAttribution string `json:"plugin_attribution"`

	// **参数解释**： 版本属性。 **约束限制**： 不涉及。 **取值范围**： - draft：草稿版本。 - formal：正式版本。 **默认取值**： 不涉及。
	VersionAttribution *string `json:"version_attribution,omitempty"`
}

func (o PluginPartQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginPartQueryDto struct{}"
	}

	return strings.Join([]string{"PluginPartQueryDto", string(data)}, " ")
}
