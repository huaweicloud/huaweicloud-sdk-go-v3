package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcEipVnic 公网IP绑定实例有PORT时，PORT的相关信息
type HwcEipVnic struct {

	// 私网IP地址
	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	// 端口所属设备ID
	DeviceId *string `json:"device_id,omitempty"`

	// 设备所属 取值范围：合法设备所属 network:dhcp network:VIP_PORT network:router_interface_distributed network:router_centralized_snat 约束：不支持设置和更新,由系统自动维护
	DeviceOwner *string `json:"device_owner,omitempty"`

	// 虚拟私有云ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 端口ID
	PortId *string `json:"port_id,omitempty"`

	// 端口profile信息
	PortProfile *string `json:"port_profile,omitempty"`

	// 端口MAC地址
	Mac *string `json:"mac,omitempty"`

	// VTEP IP
	Vtep *string `json:"vtep,omitempty"`

	// VXLAN ID
	Vni *string `json:"vni,omitempty"`

	// 端口所属实例ID,例如RDS实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 端口所属实例类型,例如“RDS”
	InstanceType *string `json:"instance_type,omitempty"`
}

func (o HwcEipVnic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcEipVnic struct{}"
	}

	return strings.Join([]string{"HwcEipVnic", string(data)}, " ")
}
