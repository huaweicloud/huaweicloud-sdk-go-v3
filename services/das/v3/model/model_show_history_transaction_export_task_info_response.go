package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHistoryTransactionExportTaskInfoResponse Response Object
type ShowHistoryTransactionExportTaskInfoResponse struct {

	// 导出任务id
	TaskId *int64 `json:"task_id,omitempty"`

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 任务状态
	TaskStatus *int32 `json:"task_status,omitempty"`

	// 导出任务开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 导出任务结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 导出任务创建时间
	CreateAt *int64 `json:"create_at,omitempty"`

	// 导出记录行数
	ExportLineNum *int64 `json:"export_line_num,omitempty"`

	// 导出文件OBS下载链接
	DownloadUrl    *string `json:"download_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHistoryTransactionExportTaskInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoryTransactionExportTaskInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowHistoryTransactionExportTaskInfoResponse", string(data)}, " ")
}
