package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlowLogExportTask 慢日志导出任务
type SlowLogExportTask struct {

	// 任务ID
	TaskId *int64 `json:"task_id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 任务状态
	TaskStatus *int32 `json:"task_status,omitempty"`

	// 任务创建时间（Unix timestamp），单位：毫秒
	CreateAt *int64 `json:"create_at,omitempty"`

	// 任务开始时间（Unix timestamp），单位：毫秒
	StartTime *int64 `json:"start_time,omitempty"`

	// 任务结束时间（Unix timestamp），单位：毫秒
	EndTime *int64 `json:"end_time,omitempty"`

	// 下载地址
	DownloadUrl *string `json:"download_url,omitempty"`
}

func (o SlowLogExportTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogExportTask struct{}"
	}

	return strings.Join([]string{"SlowLogExportTask", string(data)}, " ")
}
