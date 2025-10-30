package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerSpec 服务器规格
type ServerSpec struct {

	// 服务器规格ID
	Id *string `json:"id,omitempty"`

	// 服务器规格名称。
	Name *string `json:"name,omitempty"`

	ServerType *ServerType `json:"server_type,omitempty"`

	// 服务器发放的资源规格类型
	FlavorType *string `json:"flavor_type,omitempty"`

	// 服务器规格分类。如通用计算型/云桌面型/网关型等。
	PerformanceType *string `json:"performance_type,omitempty"`

	// 服务器功率（单位：w）
	Power *int32 `json:"power,omitempty"`

	// 设备高度。U位数
	Unit *int32 `json:"unit,omitempty"`

	// 可用虚拟CPU核数
	Vcpus *int32 `json:"vcpus,omitempty"`

	// 可用内存大小。单位：GB
	Memory *int32 `json:"memory,omitempty"`

	// 可用存储容量。单位：TiB
	StorageCapacity *int32 `json:"storage_capacity,omitempty"`

	// 名称
	CpuName *string `json:"cpu_name,omitempty"`

	// CPU架构
	CpuArchitecture *string `json:"cpu_architecture,omitempty"`
}

func (o ServerSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerSpec struct{}"
	}

	return strings.Join([]string{"ServerSpec", string(data)}, " ")
}
