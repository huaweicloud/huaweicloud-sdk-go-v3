package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiResponseDataItem struct {

	// 消费类型，具体含义根据业务定义
	ConsumptionType int32 `json:"consumption_type"`

	// 创建时间，单位为毫秒的时间戳
	CreateTime int64 `json:"create_time"`

	// 数据空间ID，唯一标识
	DataspaceId string `json:"dataspace_id"`

	// 数据空间名称
	DataspaceName string `json:"dataspace_name"`

	// 域ID，唯一标识
	DomainId string `json:"domain_id"`

	// 唯一标识ID
	Id int32 `json:"id"`

	// 管道ID，唯一标识
	PipeId string `json:"pipe_id"`

	// 管道名称
	PipeName string `json:"pipe_name"`

	// 项目ID，唯一标识
	ProjectId string `json:"project_id"`

	ShipperDestination *ShipperDestination `json:"shipper_destination"`

	// 船运ID，唯一标识
	ShipperId string `json:"shipper_id"`

	// 船运名称
	ShipperName string `json:"shipper_name"`

	ShipperSource *ShipperSource `json:"shipper_source"`

	// 状态码，具体含义根据业务定义
	Status int32 `json:"status"`

	// 更新时间，单位为毫秒的时间戳
	UpdateTime int64 `json:"update_time"`

	// 版本信息
	Version string `json:"version"`

	// 工作空间ID，唯一标识
	WorkspaceId string `json:"workspace_id"`
}

func (o ApiResponseDataItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiResponseDataItem struct{}"
	}

	return strings.Join([]string{"ApiResponseDataItem", string(data)}, " ")
}
