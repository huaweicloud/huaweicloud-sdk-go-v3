package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DbConfigCheckRequestV3 HTAP库配置校验请求体。
type DbConfigCheckRequestV3 struct {

	// GaussDB(for MySQL)实例ID。
	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	// GaussDB(for MySQL)只读节点ID。如为空，则取GaussDB(for MySQL)主节点ID
	SourceNodeId *string `json:"source_node_id,omitempty"`

	// GaussDB(for MySQL)数据库名。
	SourceDatabaseName *string `json:"source_database_name,omitempty"`

	// 目标数据库名。字符长度限制3~128位，仅支持英文大小写字母、数字以及下划线_。
	TargetDatabaseName *string `json:"target_database_name,omitempty"`

	// 同步任务名称。字符长度限制3~128位，仅支持英文大小写字母、数字以及下划线_。
	TaskName *string `json:"task_name,omitempty"`

	// 库配置列表。
	DbConfigs *[]DbConfig `json:"db_configs,omitempty"`

	TableReplConfig *TableReplConfig `json:"table_repl_config,omitempty"`
}

func (o DbConfigCheckRequestV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbConfigCheckRequestV3 struct{}"
	}

	return strings.Join([]string{"DbConfigCheckRequestV3", string(data)}, " ")
}
