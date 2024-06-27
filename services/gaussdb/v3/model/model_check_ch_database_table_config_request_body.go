package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckChDatabaseTableConfigRequestBody 表配置校验请求体。
type CheckChDatabaseTableConfigRequestBody struct {

	// 源实例ID。
	SourceInstanceId string `json:"source_instance_id"`

	// 源节点ID。GaussDB(for MySQL)只读节点ID。如为空，则取GaussDB(for MySQL)主节点ID。
	SourceNodeId *string `json:"source_node_id,omitempty"`

	// 源数据库名。
	SourceDatabaseName string `json:"source_database_name"`

	// 库配置列表。
	DbConfigs []ChDatabaseConfigsInfo `json:"db_configs"`

	// 表配置信息。
	TablesConfigs []ChDatabaseTablesConfigsInfo `json:"tables_configs"`

	TableReplConfig *ChDatabaseTableReplConfigInfo `json:"table_repl_config"`
}

func (o CheckChDatabaseTableConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckChDatabaseTableConfigRequestBody struct{}"
	}

	return strings.Join([]string{"CheckChDatabaseTableConfigRequestBody", string(data)}, " ")
}
