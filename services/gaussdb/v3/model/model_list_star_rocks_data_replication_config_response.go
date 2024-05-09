package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStarRocksDataReplicationConfigResponse Response Object
type ListStarRocksDataReplicationConfigResponse struct {

	// GaussDB(for MySQL)实例ID。
	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	// GaussDB(for MySQL)节点ID。
	SourceNodeId *string `json:"source_node_id,omitempty"`

	DatabaseInfo *DataBaseInfo `json:"database_info,omitempty"`

	// 表配置信息。
	TableInfos *[]TableConfigCheckResult `json:"table_infos,omitempty"`

	TableReplConfig *TableReplConfig `json:"table_repl_config,omitempty"`

	NewTableReplConfig *TableReplConfig `json:"new_table_repl_config,omitempty"`

	// 目标数据库名。
	TargetDatabaseName *string `json:"target_database_name,omitempty"`

	// 同步任务表是否变化。
	IsTablesChange *bool `json:"is_tables_change,omitempty"`

	// 最近一次alter table的异常信息。
	LastErrorOfAlterTable *string `json:"last_error_of_alter_table,omitempty"`
	HttpStatusCode        int     `json:"-"`
}

func (o ListStarRocksDataReplicationConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStarRocksDataReplicationConfigResponse struct{}"
	}

	return strings.Join([]string{"ListStarRocksDataReplicationConfigResponse", string(data)}, " ")
}
