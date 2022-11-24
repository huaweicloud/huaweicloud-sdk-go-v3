package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘实例地址
type InstanceAddress struct {

	// MAC地址。
	OSEXTIPSMACmacAddr *string `json:"OS-EXT-IPS-MAC:mac_addr,omitempty"`

	// IP地址对应的端口ID。
	OSEXTIPSportId *string `json:"OS-EXT-IPS:port_id,omitempty"`

	// IP地址类型。  - fixed：代表私有IP地址。 - floating：代表浮动IP地址。
	OSEXTIPStype *string `json:"OS-EXT-IPS:type,omitempty"`

	// IP地址。
	Addr *string `json:"addr,omitempty"`

	// IP地址版本。  - “4”：代表IPv4。 [- “6”：代表IPv6。](tag:hide)
	Version *string `json:"version,omitempty"`
}

func (o InstanceAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceAddress struct{}"
	}

	return strings.Join([]string{"InstanceAddress", string(data)}, " ")
}
