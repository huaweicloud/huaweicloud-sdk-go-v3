package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginTemplatesResponse Response Object
type ListPluginTemplatesResponse struct {

	// **参数解释**：API版本。 **取值范围**：可选值如下： - v1。
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**：API类型。 **取值范围**：可选值如下： - PluginTemplateList：插件模板列表。
	Kind *string `json:"kind,omitempty"`

	// **参数解释**：插件模板列表。
	Items          *[]PluginTemplateV2 `json:"items,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListPluginTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListPluginTemplatesResponse", string(data)}, " ")
}
