package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorDetailInfos flavor数据结构说明
type FlavorDetailInfos struct {

	// 裸金属服务器规格ID
	Id *string `json:"id,omitempty"`

	// 裸金属服务器规格名称
	Name *string `json:"name,omitempty"`

	// 该裸金属服务器规格对应要求系统盘大小，0为不限制。
	Disk *int32 `json:"disk,omitempty"`

	// 该裸金属服务器规格对应的CPU核数
	Vcpus *int32 `json:"vcpus,omitempty"`

	// 该裸金属服务器规格对应的内存大小，单位为MB
	Ram *int32 `json:"ram,omitempty"`

	// 该裸金属服务器规格对应的GPU设备。
	Gpus *[]GpuInfo `json:"gpus,omitempty"`

	// 该裸金属服务器规格对应的ASIC设备。
	AsicAccelerators *[]AsicAcceleratorInfo `json:"asic_accelerators,omitempty"`
}

func (o FlavorDetailInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorDetailInfos struct{}"
	}

	return strings.Join([]string{"FlavorDetailInfos", string(data)}, " ")
}
