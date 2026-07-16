package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePoolPluginResponse Response Object
type CreatePoolPluginResponse struct {

	// **参数解释**： API版本。 **取值范围**： 可选值如下： - v2
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： 插件实例的类型。 **取值范围**： 可选值如下： - Plugin：插件。
	Kind *string `json:"kind,omitempty"`

	Metadata *PluginMetadata `json:"metadata,omitempty"`

	Spec *PluginSpec `json:"spec,omitempty"`

	Status         *PluginStatus `json:"status,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreatePoolPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolPluginResponse struct{}"
	}

	return strings.Join([]string{"CreatePoolPluginResponse", string(data)}, " ")
}
