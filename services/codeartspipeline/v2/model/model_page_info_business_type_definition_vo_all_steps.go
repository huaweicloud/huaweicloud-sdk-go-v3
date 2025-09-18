package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoBusinessTypeDefinitionVoAllSteps struct {

	// **参数解释**： 插件名。 **取值范围**： 不涉及。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 展示名。 **取值范围**： 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 版本。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`
}

func (o PageInfoBusinessTypeDefinitionVoAllSteps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoBusinessTypeDefinitionVoAllSteps struct{}"
	}

	return strings.Join([]string{"PageInfoBusinessTypeDefinitionVoAllSteps", string(data)}, " ")
}
