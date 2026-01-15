package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindingVifDetails
type BindingVifDetails struct {

	// 是否为虚拟机的主网卡。
	PrimaryInterface *bool `json:"primary_interface,omitempty"`

	// 是否提供端口过滤特性, 如安全组和反MAC/IP欺骗。
	PortFilter *bool `json:"port_filter,omitempty"`

	// 是否为ovs/bridge混合模式。
	OvsHybridPlug *bool `json:"ovs_hybrid_plug,omitempty"`

	// 辅助弹性网卡的vlan ID。
	VlanId *string `json:"vlan_id,omitempty"`

	// 辅助弹性网卡的宿主网卡ID。
	ParentId *string `json:"parent_id,omitempty"`

	// 辅助弹性网卡的宿主网卡所属的设备ID。
	ParentDeviceId *string `json:"parent_device_id,omitempty"`
}

func (o BindingVifDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindingVifDetails struct{}"
	}

	return strings.Join([]string{"BindingVifDetails", string(data)}, " ")
}
