package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoolPluginsResponse Response Object
type ListPoolPluginsResponse struct {

	// **参数解释**： API版本。 **取值范围**： 可选值如下： - v2
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： 资源类型。 **取值范围**： 可选值如下： - PluginList：插件列表。
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： 资源池插件列表。
	Items          *[]Plugin `json:"items,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPoolPluginsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolPluginsResponse struct{}"
	}

	return strings.Join([]string{"ListPoolPluginsResponse", string(data)}, " ")
}
