package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListServersRspServers struct {

	// 类型，render-desktop表示渲染桌面，wdh-desktop表示专属主机桌面。
	Type *string `json:"type,omitempty"`

	// 创建时间。
	Created *string `json:"created,omitempty"`

	// 桌面id。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面套餐ID。
	DesktopProductId *string `json:"desktop_product_id,omitempty"`

	// 机器名称。
	ComputerName *string `json:"computer_name,omitempty"`

	// 机器状态。
	Status *string `json:"status,omitempty"`

	// 租户id。
	TenantId *string `json:"tenant_id,omitempty"`

	// 更新时间。
	Updated *string `json:"updated,omitempty"`

	// 任务状态。
	TaskState *string `json:"task_state,omitempty"`

	// 镜像id。
	ImageId *string `json:"image_id,omitempty"`

	// 镜像名称。
	ImageName *string `json:"image_name,omitempty"`

	// CPU核数。
	Vcpu *int32 `json:"vcpu,omitempty"`

	// 内存大小, 单位MB。
	Memory *int32 `json:"memory,omitempty"`

	// ip地址。
	IpAddresses *[]string `json:"ip_addresses,omitempty"`

	// 区域。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 桌面任务进度， 取值范围0-100以及null，null表示该桌面无任务，0-100表明该任务进度的百分比。
	Process *int32 `json:"process,omitempty"`
}

func (o ListServersRspServers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersRspServers struct{}"
	}

	return strings.Join([]string{"ListServersRspServers", string(data)}, " ")
}
