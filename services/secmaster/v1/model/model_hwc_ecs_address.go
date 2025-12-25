package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcEcsAddress 弹性云服务器的网络属性。
type HwcEcsAddress struct {

	// IP地址版本。 “4”：代表IPv4。 “6”：代表IPv6。
	Version string `json:"version"`

	// IP地址。
	Addr string `json:"addr"`

	// IP地址类型。 fixed：代表私有IP地址。 floating：代表浮动IP地址
	Type string `json:"type"`

	// MAC地址。
	MacAddr string `json:"mac_addr"`

	// IP地址对应的端口ID。
	PortId string `json:"port_id"`

	// 所属虚拟私有云的ID。
	VpcId string `json:"vpc_id"`
}

func (o HwcEcsAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcEcsAddress struct{}"
	}

	return strings.Join([]string{"HwcEcsAddress", string(data)}, " ")
}
