package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListShippersRequest Request Object
type ListShippersRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 数据空间ID
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 投递名称
	ShipperName *string `json:"shipper_name,omitempty"`

	// 数据源区域
	ShipperSourceRegion *string `json:"shipper_source_region,omitempty"`

	// 数据源消费策略
	ShipperSourceStrategy *string `json:"shipper_source_strategy,omitempty"`

	// 数据源消费类型
	ShipperConsumptionType *string `json:"shipper_consumption_type,omitempty"`

	// 目的投递类型
	DestinationShipperType *string `json:"destination_shipper_type,omitempty"`

	// 数据投递状态
	ShipperStatus *string `json:"shipper_status,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 每页数量
	Limit *int64 `json:"limit,omitempty"`

	// 第几页
	Offset *int64 `json:"offset,omitempty"`
}

func (o ListShippersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListShippersRequest struct{}"
	}

	return strings.Join([]string{"ListShippersRequest", string(data)}, " ")
}
