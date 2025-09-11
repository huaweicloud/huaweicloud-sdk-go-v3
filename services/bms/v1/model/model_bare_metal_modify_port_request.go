package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BareMetalModifyPortRequest struct {
	SubnetId *string `json:"subnet_id,omitempty"`

	DeviceOwner *string `json:"device_owner,omitempty"`

	IpAddresses *[]string `json:"ip_addresses,omitempty"`

	ReverseBinding *[]bool `json:"reverse_binding,omitempty"`
}

func (o BareMetalModifyPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BareMetalModifyPortRequest struct{}"
	}

	return strings.Join([]string{"BareMetalModifyPortRequest", string(data)}, " ")
}
