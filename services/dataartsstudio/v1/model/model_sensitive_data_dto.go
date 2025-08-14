package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SensitiveDataDto struct {

	// 唯一标识，自动生成的ID。
	Id *int64 `json:"id,omitempty"`

	// 集群名称。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 数据库名。
	DatabaseName *string `json:"database_name,omitempty"`

	// 模式名。
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名称。
	TableName *string `json:"table_name,omitempty"`

	// 字段名称。
	ColumnName *string `json:"column_name,omitempty"`

	// 规则id。
	RuleId *string `json:"rule_id,omitempty"`

	// 是否有效。1:有效 2:无效 3:待确认。
	ActiveFlag *int32 `json:"active_flag,omitempty"`

	// 数据密级id。
	ClassificationId *string `json:"classification_id,omitempty"`

	// 分类ID。
	CategoryId *string `json:"category_id,omitempty"`

	// 同步时间。
	SyncTime *int64 `json:"sync_time,omitempty"`

	// 最后发现时间。
	FindTime *int64 `json:"find_time,omitempty"`

	// 任务id。
	TaskId *string `json:"task_id,omitempty"`
}

func (o SensitiveDataDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SensitiveDataDto struct{}"
	}

	return strings.Join([]string{"SensitiveDataDto", string(data)}, " ")
}
