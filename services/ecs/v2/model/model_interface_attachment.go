package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InterfaceAttachment struct {

	// 网卡私网IP信息列表。
	FixedIps *[]ServerInterfaceFixedIp `json:"fixed_ips,omitempty" xml:"fixed_ips"`

	// 网卡Mac地址信息。
	MacAddr *string `json:"mac_addr,omitempty" xml:"mac_addr"`

	// 网卡端口所属网络ID。
	NetId *string `json:"net_id,omitempty" xml:"net_id"`

	// 网卡端口ID。
	PortId *string `json:"port_id,omitempty" xml:"port_id"`

	// 网卡端口状态。
	PortState *string `json:"port_state,omitempty" xml:"port_state"`

	// 卸载网卡时，是否删除网卡。
	DeleteOnTermination *bool `json:"delete_on_termination,omitempty" xml:"delete_on_termination"`

	// 从guest os中，网卡的驱动类型。可选值为virtio和hinic，默认为virtio
	DriverMode *string `json:"driver_mode,omitempty" xml:"driver_mode"`

	// 网卡带宽下限。
	MinRate *int32 `json:"min_rate,omitempty" xml:"min_rate"`

	// 网卡多队列个数。
	MultiqueueNum *int32 `json:"multiqueue_num,omitempty" xml:"multiqueue_num"`

	// 弹性网卡在Linux GuestOS里的BDF号
	PciAddress *string `json:"pci_address,omitempty" xml:"pci_address"`
}

func (o InterfaceAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterfaceAttachment struct{}"
	}

	return strings.Join([]string{"InterfaceAttachment", string(data)}, " ")
}
