package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BinlogExportTaskDetailResp binlog导出任务详情
type BinlogExportTaskDetailResp struct {

	// 任务ID
	TaskId *int64 `json:"task_id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 任务状态。取值范围：0（初始化）、1（运行中）、2（部分成功）、3（成功）、4（失败）、-1（已删除）
	TaskStatus *int32 `json:"task_status,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 最后记录时间
	LastRecordTime *int64 `json:"last_record_time,omitempty"`

	// 任务创建时间
	CreateAt *int64 `json:"create_at,omitempty"`

	// 导出行数
	ExportLineNum *int64 `json:"export_line_num,omitempty"`

	// 文件下载地址
	DownloadUrl *string `json:"download_url,omitempty"`

	// binlog源文件名
	SourceFileName *string `json:"source_file_name,omitempty"`

	// 解析任务ID
	ParseTaskId *int64 `json:"parse_task_id,omitempty"`
}

func (o BinlogExportTaskDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BinlogExportTaskDetailResp struct{}"
	}

	return strings.Join([]string{"BinlogExportTaskDetailResp", string(data)}, " ")
}
