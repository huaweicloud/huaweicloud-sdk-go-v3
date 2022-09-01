package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘实例地址
type InstanceAddress struct {

	// MAC地址。
	OSEXTIPSMACmacAddr *string `json:"OS-EXT-IPS-MAC:mac_addr,omitempty" xml:"OS-EXT-IPS-MAC:mac_addr"`

	// IP地址对应的端口ID。
	OSEXTIPSportId *string `json:"OS-EXT-IPS:port_id,omitempty" xml:"OS-EXT-IPS:port_id"`

	// IP地址类型。  - fixed：代表私有IP地址。 - floating：代表浮动IP地址。
	OSEXTIPStype *string `json:"OS-EXT-IPS:type,omitempty" xml:"OS-EXT-IPS:type"`

	// IP地址。
	Addr *string `json:"addr,omitempty" xml:"addr"`

	// IP地址版本。  - “4”：代表IPv4。 [- “6”：代表IPv6。](tag:hide)
	Version *string `json:"version,omitempty" xml:"version"`
}

func (o InstanceAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceAddress struct{}"
	}

	return strings.Join([]string{"InstanceAddress", string(data)}, " ")
}
