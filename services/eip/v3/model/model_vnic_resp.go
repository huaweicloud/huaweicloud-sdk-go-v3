package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VnicResp 弹性公网IP中的vnic对象，存储绑定PORT的相关信息
type VnicResp struct {

	// - 功能说明：PORT的内网地址
	PrivateIpAddress *string `json:"private_ip_address,omitempty"`

	// - 功能说明：PORT的device_id - 约束：存在PORT时，此字段associate_instance_id相同，都为实例ID
	DeviceId *string `json:"device_id,omitempty"`

	// - 功能说明：PORT的device_owner - 约束：存在PORT时，此字段和associate_instance_type都可区分实例类型
	DeviceOwner *string `json:"device_owner,omitempty"`

	// - 功能说明：PORT所在VPC的ID
	VpcId *string `json:"vpc_id,omitempty"`

	// - 功能说明：PORT的唯一标识
	PortId *string `json:"port_id,omitempty"`

	// - 功能说明：PORT的MAC信息
	Mac *string `json:"mac,omitempty"`

	// - 功能说明：PORT的使用者，不同于device_id的归属者。举例：vip port的device_owner为vip，但是这个port实际使用者可能是虚机或其他
	InstanceId *string `json:"instance_id,omitempty"`

	// - 功能说明：标记PORT使用者，与instance_id组合使用
	InstanceType *string `json:"instance_type,omitempty"`
}

func (o VnicResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VnicResp struct{}"
	}

	return strings.Join([]string{"VnicResp", string(data)}, " ")
}
