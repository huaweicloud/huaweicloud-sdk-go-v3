package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Plugin 插件实例的详细信息。
type Plugin struct {

	// **参数解释**： API版本。 **取值范围**： 可选值如下： - v2
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： 插件实例的类型。 **取值范围**： 可选值如下： - Plugin：插件。
	Kind string `json:"kind"`

	Metadata *PluginMetadata `json:"metadata"`

	Spec *PluginSpec `json:"spec"`

	Status *PluginStatus `json:"status,omitempty"`
}

func (o Plugin) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Plugin struct{}"
	}

	return strings.Join([]string{"Plugin", string(data)}, " ")
}
