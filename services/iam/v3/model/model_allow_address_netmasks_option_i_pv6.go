package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowAddressNetmasksOptionIPv6
type AllowAddressNetmasksOptionIPv6 struct {

	// IPv6地址或网段，例如： 0000:0000:0000:0000:0000:0000:0000:0000/126。
	AddressNetmask string `json:"address_netmask"`

	// 描述信息。
	Description *string `json:"description,omitempty"`
}

func (o AllowAddressNetmasksOptionIPv6) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowAddressNetmasksOptionIPv6 struct{}"
	}

	return strings.Join([]string{"AllowAddressNetmasksOptionIPv6", string(data)}, " ")
}
