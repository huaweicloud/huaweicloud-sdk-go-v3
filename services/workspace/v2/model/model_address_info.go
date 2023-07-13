package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddressInfo struct {

	// IP地址信息。
	Addr *string `json:"addr,omitempty"`

	// IP地址类型，值为4或6，分别表示IPV4和IPV6。
	Version *string `json:"version,omitempty"`

	// MAC地址。
	OSEXTIPSMACmacAddr *string `json:"OS-EXT-IPS-MAC:mac_addr,omitempty"`

	// IP地址分配方式。 -fixed  私有IP地址。 -floating 浮动IP地址。
	OSEXTIPStype *string `json:"OS-EXT-IPS:type,omitempty"`
}

func (o AddressInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddressInfo struct{}"
	}

	return strings.Join([]string{"AddressInfo", string(data)}, " ")
}
