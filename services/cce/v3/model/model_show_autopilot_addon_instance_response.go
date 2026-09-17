package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutopilotAddonInstanceResponse Response Object
type ShowAutopilotAddonInstanceResponse struct {

	// **参数解释**： API类型，固定值\"Addon\"，该值不可修改。 **约束限制**： 该值不可修改 **取值范围**： - Addon  **默认取值**： Addon
	Kind *string `json:"kind,omitempty"`

	// **参数解释**： API版本，固定值\"v3\"，该值不可修改。 **约束限制**： 该值不可修改 **取值范围**： - v3  **默认取值**： v3
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *AddonMetadata `json:"metadata,omitempty"`

	Spec *InstanceSpec `json:"spec,omitempty"`

	Status         *AddonInstanceStatus `json:"status,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowAutopilotAddonInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutopilotAddonInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowAutopilotAddonInstanceResponse", string(data)}, " ")
}
