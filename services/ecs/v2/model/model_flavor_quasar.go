package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorQuasar 云服务器规格信息。
type FlavorQuasar struct {

	// 云服务器类型ID。
	Id *string `json:"id,omitempty"`

	// 云服务器规格名称。
	Name *string `json:"name,omitempty"`

	// 该云服务器规格对应的CPU核数。
	Vcpus *int32 `json:"vcpus,omitempty"`

	// 该云服务器规格对应的内存大小，单位为MB.
	Ram *int32 `json:"ram,omitempty"`

	// 该云服务器规格对应要求系统盘大小，0为不限制。
	Disk *int32 `json:"disk,omitempty"`

	Gpus *[]GpuInfo `json:"gpus,omitempty"`

	AsicAccelerators *[]AsicAcceleratorInfo `json:"asic_accelerators,omitempty"`
}

func (o FlavorQuasar) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorQuasar struct{}"
	}

	return strings.Join([]string{"FlavorQuasar", string(data)}, " ")
}
