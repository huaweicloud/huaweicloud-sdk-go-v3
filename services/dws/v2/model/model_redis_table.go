package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RedisTable 重分布表信息
type RedisTable struct {

	// 表名
	TableName *string `json:"table_name,omitempty"`

	// 表唯一id
	Id *int32 `json:"id,omitempty"`

	// schema名
	SchemaName *string `json:"schema_name,omitempty"`

	// 逻辑集群名
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 表大小
	Size *int64 `json:"size,omitempty"`

	// 重分布类型 i：重分布中； y：重分布完成； n：未开始
	Status *string `json:"status,omitempty"`
}

func (o RedisTable) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedisTable struct{}"
	}

	return strings.Join([]string{"RedisTable", string(data)}, " ")
}
