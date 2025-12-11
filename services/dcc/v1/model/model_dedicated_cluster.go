package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DedicatedCluster struct {

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 可用区
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 主机数量
	HostTotal *int32 `json:"host_total,omitempty"`

	// 主机规格编码。
	HostType *string `json:"host_type,omitempty"`

	// 集群服务类型。
	ServiceType *string `json:"service_type,omitempty"`

	HostProperties *HostProperties `json:"host_properties,omitempty"`

	// 已用vcpu个数。
	VcpusUsed *int32 `json:"vcpus_used,omitempty"`

	// 总的vcpu个数。。
	VcpusTotal *int32 `json:"vcpus_total,omitempty"`

	// 已用内存。
	MemoryMbUsed *int32 `json:"memory_mb_used,omitempty"`

	// 总内存。
	MemoryMbTotal *int32 `json:"memory_mb_total,omitempty"`

	// 支持flavor列表
	Flavors *[]string `json:"flavors,omitempty"`

	// 运行的计算实例总数。
	InstanceTotal *int32 `json:"instance_total,omitempty"`
}

func (o DedicatedCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DedicatedCluster struct{}"
	}

	return strings.Join([]string{"DedicatedCluster", string(data)}, " ")
}
