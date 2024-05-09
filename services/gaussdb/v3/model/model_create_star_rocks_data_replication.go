package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStarRocksDataReplication StarRocks创建数据同步请求。
type CreateStarRocksDataReplication struct {

	// GaussDB(for MySQL)实例ID。
	SourceInstanceId string `json:"source_instance_id"`

	// GaussDB(for MySQL)只读节点ID。如为空，则取GaussDB(for MySQL)主节点ID
	SourceNodeId *string `json:"source_node_id,omitempty"`

	// 源数据库。
	SourceDatabase string `json:"source_database"`

	// 目标数据库。 字符长度限制3~128位，仅支持英文大小写字母、数字以及下划线_。
	TargetDatabase string `json:"target_database"`

	// 同步任务名。 字符长度限制3~128位，仅支持英文大小写字母、数字以及下划线_。
	TaskName string `json:"task_name"`
}

func (o CreateStarRocksDataReplication) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStarRocksDataReplication struct{}"
	}

	return strings.Join([]string{"CreateStarRocksDataReplication", string(data)}, " ")
}
