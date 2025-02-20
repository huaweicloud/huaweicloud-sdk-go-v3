package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobDetailRespRepairProgressInfoRepairProgressDetails 修复进度明细
type JobDetailRespRepairProgressInfoRepairProgressDetails struct {

	// 对比任务ID。
	QueryId *string `json:"query_id,omitempty"`

	// 库名。
	DbName *string `json:"db_name,omitempty"`

	// schema名。
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名。
	TableName *string `json:"table_name,omitempty"`

	// 总行数。
	TotalRowCount *int32 `json:"total_row_count,omitempty"`

	// 完成行数。
	CompleteRowCount *int32 `json:"complete_row_count,omitempty"`

	// 过滤行数。
	FilterRowCount *int32 `json:"filter_row_count,omitempty"`

	// 已修复行数。
	RepairedRowCount *int32 `json:"repaired_row_count,omitempty"`

	// 失败行数。
	FailedRowCount *int32 `json:"failed_row_count,omitempty"`

	// 修复状态。
	RepairStatus *int32 `json:"repair_status,omitempty"`

	// 开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 更新时间。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o JobDetailRespRepairProgressInfoRepairProgressDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobDetailRespRepairProgressInfoRepairProgressDetails struct{}"
	}

	return strings.Join([]string{"JobDetailRespRepairProgressInfoRepairProgressDetails", string(data)}, " ")
}
