package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateShipperBody struct {

	// 消费类型，0-实时 1-调度
	ConsumptionType *int32 `json:"consumption_type,omitempty"`

	// 数据空间ID
	DataspaceId *string `json:"dataspace_id,omitempty"`

	// 数据空间名称
	DataspaceName *string `json:"dataspace_name,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 管道ID
	PipeId *string `json:"pipe_id,omitempty"`

	// 管道名称
	PipeName *string `json:"pipe_name,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	ShipperDestination *CreateShipperBodyShipperDestination `json:"shipper_destination,omitempty"`

	// 投递名称
	ShipperName *string `json:"shipper_name,omitempty"`

	ShipperSource *CreateShipperBodyShipperSource `json:"shipper_source,omitempty"`

	// 版本信息
	Version *string `json:"version,omitempty"`

	// 工作空间ID
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 工作空间名称
	WorkspaceName *string `json:"workspace_name,omitempty"`
}

func (o CreateShipperBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShipperBody struct{}"
	}

	return strings.Join([]string{"CreateShipperBody", string(data)}, " ")
}
