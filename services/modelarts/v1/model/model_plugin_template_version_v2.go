package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginTemplateVersionV2 插件模板的版本信息。
type PluginTemplateVersionV2 struct {

	// **参数解释**：插件模板的版本号。 **取值范围**：不涉及。
	Version string `json:"version"`

	// **参数解释**：创建时间。 **取值范围**：不涉及。
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// **参数解释**：插件安装参数。
	Inputs *interface{} `json:"inputs,omitempty"`

	// **参数解释**：供界面使用的翻译信息。
	Translate *interface{} `json:"translate,omitempty"`

	// **参数解释**：版本描述信息。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：版本描述信息。 **取值范围**：不涉及。
	Detail *string `json:"detail,omitempty"`
}

func (o PluginTemplateVersionV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginTemplateVersionV2 struct{}"
	}

	return strings.Join([]string{"PluginTemplateVersionV2", string(data)}, " ")
}
