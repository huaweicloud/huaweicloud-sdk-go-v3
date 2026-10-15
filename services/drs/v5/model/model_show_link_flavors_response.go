package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLinkFlavorsResponse Response Object
type ShowLinkFlavorsResponse struct {
	VmFlavor *FlavorInfo `json:"vm_flavor,omitempty"`

	VolumeFlavor *FlavorInfo `json:"volume_flavor,omitempty"`

	FlowFlavor     *FlavorInfo `json:"flow_flavor,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowLinkFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLinkFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ShowLinkFlavorsResponse", string(data)}, " ")
}
