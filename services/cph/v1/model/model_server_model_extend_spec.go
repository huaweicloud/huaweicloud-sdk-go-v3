package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 云手机服务器的属性描述，不超过512个字节
type ServerModelExtendSpec struct {

	// 云手机服务器cpu类型
	Cpu *string `json:"cpu,omitempty"`

	// 云手机服务器内存类型
	Memory *string `json:"memory,omitempty"`

	// 云手机服务器磁盘类型
	Disk *string `json:"disk,omitempty"`

	// 云手机服务器网络类型
	NetworkInterface *string `json:"network_interface,omitempty"`

	// 云手机服务器gpu类型
	Gpu *string `json:"gpu,omitempty"`

	// 云手机服务器bms规格
	BmsFlavor *string `json:"bms_flavor,omitempty"`

	// 云手机服务器gpu数量
	GpuCount *int32 `json:"gpu_count,omitempty"`

	// 云手机服务器numa数量
	NumaCount *int32 `json:"numa_count,omitempty"`
}

func (o ServerModelExtendSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerModelExtendSpec struct{}"
	}

	return strings.Join([]string{"ServerModelExtendSpec", string(data)}, " ")
}
