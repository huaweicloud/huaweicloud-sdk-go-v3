package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShipperBodyShipperSource 投递源
type CreateShipperBodyShipperSource struct {

	// 区域信息
	Region *string `json:"region,omitempty"`

	// 源数据空间ID
	SourceDataspace *string `json:"source_dataspace,omitempty"`

	// 源数据空间名称
	SourceDataspaceName *string `json:"source_dataspace_name,omitempty"`

	// 源身份角色
	SourceIdentityRole *string `json:"source_identity_role,omitempty"`

	// 源表ID
	SourceTable *string `json:"source_table,omitempty"`

	// 源表名称
	SourceTableName *string `json:"source_table_name,omitempty"`

	// 源管道ID
	SourcePipe *string `json:"source_pipe,omitempty"`

	// 源管道名称
	SourcePipeName *string `json:"source_pipe_name,omitempty"`

	// 源类型
	SourceType *int32 `json:"source_type,omitempty"`

	// 源工作空间ID
	SourceWorkspace *string `json:"source_workspace,omitempty"`

	// 源工作空间名称
	SourceWorkspaceName *string `json:"source_workspace_name,omitempty"`
}

func (o CreateShipperBodyShipperSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShipperBodyShipperSource struct{}"
	}

	return strings.Join([]string{"CreateShipperBodyShipperSource", string(data)}, " ")
}
