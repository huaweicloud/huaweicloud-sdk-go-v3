package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginTemplateResponse Response Object
type ShowPluginTemplateResponse struct {

	// **参数解释**：API版本。 **取值范围**：可选值如下： - v1。
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**：资源类型。 **取值范围**：可选值如下： - PluginTemplate：插件模板
	Kind *string `json:"kind,omitempty"`

	Metadata *PluginTemplateMetadata `json:"metadata,omitempty"`

	Spec           *PluginTemplateSpec `json:"spec,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowPluginTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowPluginTemplateResponse", string(data)}, " ")
}
