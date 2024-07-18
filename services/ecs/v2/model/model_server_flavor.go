package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerFlavor 弹性云服务器规格信息。
type ServerFlavor struct {

	// 弹性云服务器规格ID。
	Id string `json:"id"`

	// 弹性云服务器规格名称。
	Name string `json:"name"`

	// 该云服务器规格对应要求系统盘大小，0为不限制。此字段在本系统中无效。
	Disk string `json:"disk"`

	// 该云服务器规格对应的CPU核数。
	Vcpus string `json:"vcpus"`

	// 该云服务器规格对应的内存大小，单位为MB。
	Ram string `json:"ram"`

	// 该云服务器规格对应的GPU设备。
	Gpus []GpuInfo `json:"gpus"`

	// 该云服务器规格对应的ASIC设备。
	AsicAccelerators []AsicAcceleratorInfo `json:"asic_accelerators"`
}

func (o ServerFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerFlavor struct{}"
	}

	return strings.Join([]string{"ServerFlavor", string(data)}, " ")
}
