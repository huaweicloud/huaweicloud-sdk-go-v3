package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAddonInstanceResponse Response Object
type ShowAddonInstanceResponse struct {

	// API类型，必须是\"Addon\"
	Kind *string `json:"kind,omitempty"`

	// API版本，必须是\"v3\"
	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ObjectMeta `json:"metadata,omitempty"`

	Spec *AddonInstanceSpec `json:"spec,omitempty"`

	Status         *AddonInstanceStatus `json:"status,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowAddonInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAddonInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowAddonInstanceResponse", string(data)}, " ")
}
