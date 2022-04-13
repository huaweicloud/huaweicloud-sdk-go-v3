package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AllowAddressNetmasksResult struct {
	// IP地址或网段，例如：192.168.0.1/24。

	AddressNetmask string `json:"address_netmask"`
	// 描述信息。

	Description string `json:"description"`
}

func (o AllowAddressNetmasksResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowAddressNetmasksResult struct{}"
	}

	return strings.Join([]string{"AllowAddressNetmasksResult", string(data)}, " ")
}
