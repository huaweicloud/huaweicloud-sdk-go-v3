package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowAddressNetmasksResultIPv6
type AllowAddressNetmasksResultIPv6 struct {

	// IPv6地址或网段，例如： 0000:0000:0000:0000:0000:0000:0000:0000/126。
	AddressNetmask string `json:"address_netmask"`

	// 描述信息。
	Description *string `json:"description,omitempty"`
}

func (o AllowAddressNetmasksResultIPv6) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowAddressNetmasksResultIPv6 struct{}"
	}

	return strings.Join([]string{"AllowAddressNetmasksResultIPv6", string(data)}, " ")
}
