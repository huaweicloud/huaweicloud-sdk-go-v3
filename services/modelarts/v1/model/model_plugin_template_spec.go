package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginTemplateSpec 插件模板描述信息。
type PluginTemplateSpec struct {

	// **参数解释**：插件模板类型。 **取值范围**：可选值如下： - npu-river：NPU驱动 - gpu-driver：GPU驱动
	Type string `json:"type"`

	// **参数解释**：插件模板描述。 **取值范围**：不涉及。
	Description string `json:"description"`

	// **参数解释**：插件模板版本描述信息。
	Versions map[string]PluginTemplateVersionV2 `json:"versions"`
}

func (o PluginTemplateSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginTemplateSpec struct{}"
	}

	return strings.Join([]string{"PluginTemplateSpec", string(data)}, " ")
}
