package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostsRspHostProperties 云办公主机的属性。
type ListHostsRspHostProperties struct {

	// 云办公主机的vCPUs个数。
	Vcpus *int32 `json:"vcpus,omitempty"`

	// 云办公主机的物理核数。
	Cores *int32 `json:"cores,omitempty"`

	// 云办公主机的物理套接字数量。
	Sockets *int32 `json:"sockets,omitempty"`

	// 云办公主机的物理内存大小。
	Memory *int32 `json:"memory,omitempty"`

	// 云办公主机类型。
	HostType *string `json:"host_type,omitempty"`

	// 云办公主机类型名称。
	HostTypeName *string `json:"host_type_name,omitempty"`

	// 可以创建的规格。
	AvailableInstanceCapacities *[]ListHostsRspHostPropertiesAvailableInstanceCapacities `json:"available_instance_capacities,omitempty"`
}

func (o ListHostsRspHostProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsRspHostProperties struct{}"
	}

	return strings.Join([]string{"ListHostsRspHostProperties", string(data)}, " ")
}
