package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginTemplateMetadata 插件模板的metadata信息。
type PluginTemplateMetadata struct {

	// **参数解释**：插件模板的名称。 **取值范围**：不涉及。
	Name string `json:"name"`

	// **参数解释**：插件模板注解，由key/value组成。
	Annotations map[string]string `json:"annotations,omitempty"`
}

func (o PluginTemplateMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginTemplateMetadata struct{}"
	}

	return strings.Join([]string{"PluginTemplateMetadata", string(data)}, " ")
}
