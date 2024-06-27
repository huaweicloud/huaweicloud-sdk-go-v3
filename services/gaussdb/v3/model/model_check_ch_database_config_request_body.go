package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckChDatabaseConfigRequestBody 库配置校验参数。
type CheckChDatabaseConfigRequestBody struct {

	// 源实例ID，严格匹配UUID规则。
	SourceInstanceId string `json:"source_instance_id"`

	// 源节点ID。GaussDB(for MySQL)只读节点ID。如为空，则取GaussDB(for MySQL)主节点ID。
	SourceNodeId *string `json:"source_node_id,omitempty"`

	// 源数据库名。
	SourceDatabaseName string `json:"source_database_name"`

	// 库配置列表。
	DbConfigs []ChDatabaseConfigsInfo `json:"db_configs"`

	TableReplConfig *ChDatabaseTableReplConfigInfo `json:"table_repl_config"`
}

func (o CheckChDatabaseConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckChDatabaseConfigRequestBody struct{}"
	}

	return strings.Join([]string{"CheckChDatabaseConfigRequestBody", string(data)}, " ")
}
