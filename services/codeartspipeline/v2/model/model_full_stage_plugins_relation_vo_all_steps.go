package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FullStagePluginsRelationVoAllSteps struct {

	// **参数解释**： 扩展插件插件名。 **取值范围**： 不涉及。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 扩展插件展示名。 **取值范围**： 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 扩展插件版本号。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`
}

func (o FullStagePluginsRelationVoAllSteps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullStagePluginsRelationVoAllSteps struct{}"
	}

	return strings.Join([]string{"FullStagePluginsRelationVoAllSteps", string(data)}, " ")
}
