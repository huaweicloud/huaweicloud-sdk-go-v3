package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginTemplateV2 插件模板的详细信息。
type PluginTemplateV2 struct {

	// **参数解释**：API版本。 **取值范围**：可选值如下： - v2
	ApiVersion string `json:"apiVersion"`

	// **参数解释**：资源类型。 **取值范围**：可选值如下： - PluginTemplate：插件模板
	Kind string `json:"kind"`

	Metadata *PluginTemplateMetadata `json:"metadata"`

	Spec *PluginTemplateSpecV2 `json:"spec"`
}

func (o PluginTemplateV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginTemplateV2 struct{}"
	}

	return strings.Join([]string{"PluginTemplateV2", string(data)}, " ")
}
