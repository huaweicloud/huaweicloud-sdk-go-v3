package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceVnicResp 实例port的信息
type InstanceVnicResp struct {

	// 实例port的ip地址
	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	// port的device_id
	DeviceId *string `json:"device_id,omitempty"`

	// port的device_owner
	DeviceOwner *string `json:"device_owner,omitempty"`

	// port的vpc_id
	VpcId *string `json:"vpc_id,omitempty"`

	// port的uuid
	PortId *string `json:"port_id,omitempty"`

	// port的mac地址
	Mac *string `json:"mac,omitempty"`

	// port的vtep地址
	Vtep *string `json:"vtep,omitempty"`

	// port的vni
	Vni *string `json:"vni,omitempty"`

	// port的实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// port的实例类型
	InstanceType *string `json:"instance_type,omitempty"`

	// port的profile
	PortProfile *string `json:"port_profile,omitempty"`
}

func (o InstanceVnicResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceVnicResp struct{}"
	}

	return strings.Join([]string{"InstanceVnicResp", string(data)}, " ")
}
