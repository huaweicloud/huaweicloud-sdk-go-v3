package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBinlogExportTaskInfoResponse Response Object
type ShowBinlogExportTaskInfoResponse struct {

	// 导出任务ID
	TaskId *int64 `json:"task_id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 任务状态。取值范围：0（初始化）、1（运行中）、2（部分成功）、3（成功）、4（失败）、-1（已删除）
	TaskStatus *int32 `json:"task_status,omitempty"`

	// 任务开始时间，单位毫秒
	StartTime *int64 `json:"start_time,omitempty"`

	// 任务结束时间，单位毫秒
	EndTime *int64 `json:"end_time,omitempty"`

	// 任务创建时间，单位毫秒
	CreateAt *int64 `json:"create_at,omitempty"`

	// 导出行数
	ExportLineNum *int64 `json:"export_line_num,omitempty"`

	// 导出文件下载地址
	DownloadUrl    *string `json:"download_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBinlogExportTaskInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBinlogExportTaskInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowBinlogExportTaskInfoResponse", string(data)}, " ")
}
