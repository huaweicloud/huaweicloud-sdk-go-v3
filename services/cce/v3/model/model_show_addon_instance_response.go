package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAddonInstanceResponse struct {

	// API类型，固定值“Addon”，该值不可修改。
	Kind *string `json:"kind,omitempty" xml:"kind"`

	// API版本，固定值“v3”，该值不可修改。
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion"`

	Metadata *Metadata `json:"metadata,omitempty" xml:"metadata"`

	Spec *InstanceSpec `json:"spec,omitempty" xml:"spec"`

	Status         *AddonInstanceStatus `json:"status,omitempty" xml:"status"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowAddonInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAddonInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowAddonInstanceResponse", string(data)}, " ")
}
