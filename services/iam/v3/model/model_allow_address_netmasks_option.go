package model

import (
	"encoding/json"

	"strings"
)

//
type AllowAddressNetmasksOption struct {
	// IP地址或网段,例如:192.168.0.1/24。

	AddressNetmask string `json:"address_netmask"`
	// 描述信息。

	Description *string `json:"description,omitempty"`
}

func (o AllowAddressNetmasksOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AllowAddressNetmasksOption struct{}"
	}

	return strings.Join([]string{"AllowAddressNetmasksOption", string(data)}, " ")
}
