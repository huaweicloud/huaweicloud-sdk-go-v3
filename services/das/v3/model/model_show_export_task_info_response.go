package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExportTaskInfoResponse Response Object
type ShowExportTaskInfoResponse struct {

	// 创建时间
	CreateAt float32 `json:"create_at,omitempty"`

	// 下载链接
	DownloadUrl *string `json:"download_url,omitempty"`

	// 结束时间
	EndTime float32 `json:"end_time,omitempty"`

	// 导出条数
	ExportLineNum float32 `json:"export_line_num,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 最新SQL执行时间
	LastRecordTime float32 `json:"last_record_time,omitempty"`

	// 开始时间
	StartTime float32 `json:"start_time,omitempty"`

	// 任务ID
	TaskId float32 `json:"task_id,omitempty"`

	// 任务状态
	TaskStatus     *int32 `json:"task_status,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowExportTaskInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExportTaskInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowExportTaskInfoResponse", string(data)}, " ")
}
