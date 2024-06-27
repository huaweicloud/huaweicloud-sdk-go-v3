package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClickHouseDataBaseReplicationConfigResponse Response Object
type ListClickHouseDataBaseReplicationConfigResponse struct {

	// 源实例ID。
	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	// 源实例节点ID。
	SourceNodeId *string `json:"source_node_id,omitempty"`

	DatabaseInfo *ChDatabaseConfigResponse `json:"database_info,omitempty"`

	// 表配置信息。
	TableInfos *[]ChDatabaseTableConfigCheckResult `json:"table_infos,omitempty"`

	TableReplConfig *ChDatabaseTableReplConfigInfo `json:"table_repl_config,omitempty"`

	// 同步任务表是否变化。
	IsTablesChange *bool `json:"is_tables_change,omitempty"`

	NewTableReplConfig *ChDatabaseTableReplConfigInfo `json:"new_table_repl_config,omitempty"`

	// 最近一次alter table的异常信息。
	LastErrorOfAlterTable *string `json:"last_error_of_alter_table,omitempty"`
	HttpStatusCode        int     `json:"-"`
}

func (o ListClickHouseDataBaseReplicationConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClickHouseDataBaseReplicationConfigResponse struct{}"
	}

	return strings.Join([]string{"ListClickHouseDataBaseReplicationConfigResponse", string(data)}, " ")
}
