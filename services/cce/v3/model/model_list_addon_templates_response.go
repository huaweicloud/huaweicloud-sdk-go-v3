package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAddonTemplatesResponse Response Object
type ListAddonTemplatesResponse struct {

	// **参数解释**： API类型，固定值\"Addon\"，该值不可修改。 **约束限制**： 该值不可修改 **取值范围**： - Addon  **默认取值**： Addon
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API版本，固定值\"v3\"，该值不可修改。 **约束限制**： 该值不可修改 **取值范围**： - v3  **默认取值**： v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	// **参数解释**： 插件模板列表。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Items          *[]AddonTemplate `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListAddonTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddonTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListAddonTemplatesResponse", string(data)}, " ")
}
