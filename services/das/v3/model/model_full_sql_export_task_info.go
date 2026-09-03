package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FullSqlExportTaskInfo 全量SQL导出任务信息
type FullSqlExportTaskInfo struct {

	// 任务ID
	TaskId *int64 `json:"task_id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 任务状态。取值范围：0（等待）、1（运行中）、2（失败）、3（成功）、4（超时）、5（OBS文件已删除）
	TaskStatus *int32 `json:"task_status,omitempty"`

	// 导出任务开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 导出任务结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 最后一条SQL执行的时间点
	LastRecordTime *int64 `json:"last_record_time,omitempty"`

	// 导出任务创建时间
	CreateAt *int64 `json:"create_at,omitempty"`

	// 导出行数
	ExportLineNum *int64 `json:"export_line_num,omitempty"`

	// 导出文件下载URL
	DownloadUrl *string `json:"download_url,omitempty"`
}

func (o FullSqlExportTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullSqlExportTaskInfo struct{}"
	}

	return strings.Join([]string{"FullSqlExportTaskInfo", string(data)}, " ")
}
