package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RespHostProperty 专属主机属性。
type RespHostProperty struct {

	// 专属主机类型。
	HostType string `json:"host_type"`

	// 专属主机类型的名称。
	HostTypeName string `json:"host_type_name"`

	// 专属主机的vCPUs个数。
	Vcpus int32 `json:"vcpus"`

	// 专属主机的物理核数。
	Cores int32 `json:"cores"`

	// 专属主机的物理套接字数量。
	Sockets int32 `json:"sockets"`

	// 专属主机的物理内存大小。
	Memory int32 `json:"memory"`

	// 专属主机上创建的云服务器规格。
	AvailableInstanceCapacities []RespInstanceCapacity `json:"available_instance_capacities"`
}

func (o RespHostProperty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespHostProperty struct{}"
	}

	return strings.Join([]string{"RespHostProperty", string(data)}, " ")
}
