package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostModel struct {

	// 虚拟机ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 虚拟机名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 虚拟机IP地址
	Ip *string `json:"ip,omitempty" xml:"ip"`

	// 可用区域
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty" xml:"availability_zone_id"`

	// 标签列表信息
	Tags *[]TagPlain `json:"tags,omitempty" xml:"tags"`

	// 虚拟机当前状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 虚拟机规格ID
	Flavor *string `json:"flavor,omitempty" xml:"flavor"`

	// 虚拟机类型，当前支持MasterNode，CoreNode，TaskNode
	Type *string `json:"type,omitempty" xml:"type"`

	// 内存
	Mem *string `json:"mem,omitempty" xml:"mem"`

	// CPU核数
	Cpu *string `json:"cpu,omitempty" xml:"cpu"`

	// 操作系统盘容量
	RootVolumeSize *string `json:"root_volume_size,omitempty" xml:"root_volume_size"`

	// 数据盘类型
	DataVolumeType *string `json:"data_volume_type,omitempty" xml:"data_volume_type"`

	// 数据盘容量
	DataVolumeSize *int32 `json:"data_volume_size,omitempty" xml:"data_volume_size"`

	// 数据盘个数
	DataVolumeCount *int32 `json:"data_volume_count,omitempty" xml:"data_volume_count"`
}

func (o HostModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostModel struct{}"
	}

	return strings.Join([]string{"HostModel", string(data)}, " ")
}
