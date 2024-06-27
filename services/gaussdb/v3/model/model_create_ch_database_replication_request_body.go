package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChDatabaseReplicationRequestBody ClickHouse创建数据同步请求。
type CreateChDatabaseReplicationRequestBody struct {

	// 源实例ID。
	SourceInstanceId string `json:"source_instance_id"`

	// 源节点ID。GaussDB(for MySQL)只读节点ID。如为空，则取GaussDB(for MySQL)主节点ID。
	SourceNodeId *string `json:"source_node_id,omitempty"`

	// 源数据库。
	SourceDatabase string `json:"source_database"`
}

func (o CreateChDatabaseReplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChDatabaseReplicationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateChDatabaseReplicationRequestBody", string(data)}, " ")
}
