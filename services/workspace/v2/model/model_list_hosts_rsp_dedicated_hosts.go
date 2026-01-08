package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListHostsRspDedicatedHosts struct {

	// 云办公主机ID。
	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`

	// 云办公主机的名称。
	Name *string `json:"name,omitempty"`

	// 在创建云服务器时（未指定专属主机ID），是否允许云服务器自动分配在一台可用的云办公主机上。取值范围：“on”或“off”。
	AutoPlacement *string `json:"auto_placement,omitempty"`

	// 云办公主机的区域。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	HostProperties *ListHostsRspHostProperties `json:"host_properties,omitempty"`

	// 云办公主机的产品id。
	ProductId *string `json:"product_id,omitempty"`

	// 云办公主机的订单id。
	OrderId *string `json:"order_id,omitempty"`

	// 云服务资源类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 云办公主机状态，该参数取值可以为：“available”、“fault”或“released”。
	State *string `json:"state,omitempty"`

	// 云办公主机可用的vCPU核数。
	AvailableVcpus *int32 `json:"available_vcpus,omitempty"`

	// 云办公主机可用的内存大小。
	AvailableMemory *int32 `json:"available_memory,omitempty"`

	// 云办公主机上的实例总数。
	InstanceTotal *int32 `json:"instance_total,omitempty"`

	// 云办公主机的分配时间。
	AllocatedAt *string `json:"allocated_at,omitempty"`

	// 云办公主机的释放时间。
	ReleasedAt *string `json:"released_at,omitempty"`

	// 专属主机上的实例UUID。
	InstanceUuids *[]string `json:"instance_uuids,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`
}

func (o ListHostsRspDedicatedHosts) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsRspDedicatedHosts struct{}"
	}

	return strings.Join([]string{"ListHostsRspDedicatedHosts", string(data)}, " ")
}
