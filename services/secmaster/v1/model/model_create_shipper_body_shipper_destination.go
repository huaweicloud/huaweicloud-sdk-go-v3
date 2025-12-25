package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShipperBodyShipperDestination 投递目标
type CreateShipperBodyShipperDestination struct {

	// 数据参数
	DataParam *string `json:"data_param,omitempty"`

	// 目标数据空间ID
	DestinationDataspace *string `json:"destination_dataspace,omitempty"`

	// 目标数据空间名称
	DestinationDataspaceName *string `json:"destination_dataspace_name,omitempty"`

	// 目标身份角色
	DestinationIdentityRole *string `json:"destination_identity_role,omitempty"`

	// 目标表ID
	DestinationTable *string `json:"destination_table,omitempty"`

	// 目标表名称
	DestinationTableName *string `json:"destination_table_name,omitempty"`

	// 目标管道ID
	DestinationPipe *string `json:"destination_pipe,omitempty"`

	// 目标管道名称
	DestinationPipeName *string `json:"destination_pipe_name,omitempty"`

	// 目标区域信息
	DestinationRegion *string `json:"destination_region,omitempty"`

	// 目标投递类型
	DestinationShipperType *int32 `json:"destination_shipper_type,omitempty"`

	// 目标工作空间ID
	DestinationWorkspace *string `json:"destination_workspace,omitempty"`

	// 目标工作空间名称
	DestinationWorkspaceName *string `json:"destination_workspace_name,omitempty"`
}

func (o CreateShipperBodyShipperDestination) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShipperBodyShipperDestination struct{}"
	}

	return strings.Join([]string{"CreateShipperBodyShipperDestination", string(data)}, " ")
}
