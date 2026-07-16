package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginMetadata 插件实例的metadata信息。
type PluginMetadata struct {

	// **参数解释**： 插件实例的名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`
}

func (o PluginMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginMetadata struct{}"
	}

	return strings.Join([]string{"PluginMetadata", string(data)}, " ")
}
