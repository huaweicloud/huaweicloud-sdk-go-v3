package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostModel struct {
	// 虚拟机ID

	Id *string `json:"id,omitempty"`
	// 虚拟机名称

	Name *string `json:"name,omitempty"`
	// 虚拟机IP地址

	Ip *string `json:"ip,omitempty"`
	// 虚拟机当前状态

	Status *string `json:"status,omitempty"`
	// 虚拟机规格ID

	Flavor *string `json:"flavor,omitempty"`
	// 虚拟机类型，当前支持MasterNode，CoreNode，TaskNode

	Type *string `json:"type,omitempty"`
	// 内存

	Mem *string `json:"mem,omitempty"`
	// CPU核数

	Cpu *string `json:"cpu,omitempty"`
	// 操作系统盘容量

	RootVolumeSize *string `json:"root_volume_size,omitempty"`
	// 数据盘类型

	DataVolumeType *string `json:"data_volume_type,omitempty"`
	// 数据盘容量

	DataVolumeSize *int32 `json:"data_volume_size,omitempty"`
	// 数据盘个数

	DataVolumeCount *int32 `json:"data_volume_count,omitempty"`
}

func (o HostModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostModel struct{}"
	}

	return strings.Join([]string{"HostModel", string(data)}, " ")
}
