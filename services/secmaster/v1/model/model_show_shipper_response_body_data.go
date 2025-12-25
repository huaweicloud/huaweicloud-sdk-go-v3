package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShipperResponseBodyData 投递数据
type ShowShipperResponseBodyData struct {

	// 消费类型
	ConsumptionType *int32 `json:"consumption_type,omitempty"`

	// 创建时间（时间戳）
	CreateTime *int64 `json:"create_time,omitempty"`

	// 数据空间ID
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 数据空间名称
	DataspaceName *string `json:"dataspace_name,omitempty"`

	// 域ID
	DomainId *string `json:"domain_id,omitempty"`

	// 管道ID
	Id *int32 `json:"id,omitempty"`

	// 管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 管道名称
	PipeName *string `json:"pipe_name,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	ShipperDestination *ShowShipperResponseBodyDataShipperDestination `json:"shipper_destination,omitempty"`

	// 投递ID
	ShipperId *string `json:"shipper_id,omitempty"`

	// 投递名称
	ShipperName *string `json:"shipper_name,omitempty"`

	ShipperSource *ShowShipperResponseBodyDataShipperSource `json:"shipper_source,omitempty"`

	// 状态
	Status *int32 `json:"status,omitempty"`

	// 更新时间（时间戳）
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 工作空间名称
	WorkspaceName *string `json:"workspace_name,omitempty"`
}

func (o ShowShipperResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShipperResponseBodyData struct{}"
	}

	return strings.Join([]string{"ShowShipperResponseBodyData", string(data)}, " ")
}
