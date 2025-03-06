package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkspaceTransformVo struct {

	// 所属关系建模的模型ID，ID字符串。
	TargetModelId *string `json:"target_model_id,omitempty"`

	// 工作区名字。
	TargetModelName string `json:"target_model_name"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 是否更新已有表。
	UpdateExistTables *bool `json:"update_exist_tables,omitempty"`

	// 需要物化的逻辑实体的ID列表，ID字符串。
	Ids *[]string `json:"ids,omitempty"`

	// 数据连接类型，对应表所在的数仓类型，取值可以为DLI、DWS、MRS_HIVE、POSTGRESQL、MRS_SPARK、CLICKHOUSE、MYSQL、ORACLE和DORIS等。
	DwType string `json:"dw_type"`

	// 转化后物理表所属的数据连接ID。
	ConnectionId *string `json:"connection_id,omitempty"`

	// 转化后物理表所属的数据库。
	Database *string `json:"database,omitempty"`

	// 转化后物理表所属的队列（仅DLI）。
	Queue *string `json:"queue,omitempty"`

	// 转化后物理表所属的schema（仅DWS和PostgreSQL）。
	Schema *string `json:"schema,omitempty"`
}

func (o WorkspaceTransformVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkspaceTransformVo struct{}"
	}

	return strings.Join([]string{"WorkspaceTransformVo", string(data)}, " ")
}
