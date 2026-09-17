package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndexUsageExportTaskInfo 索引使用导出任务信息
type IndexUsageExportTaskInfo struct {

	// 任务ID
	TaskId *int64 `json:"task_id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 任务状态
	TaskStatus *int32 `json:"task_status,omitempty"`

	// 创建时间
	CreateAt *int64 `json:"create_at,omitempty"`

	// 下载地址
	DownloadUrl *string `json:"download_url,omitempty"`
}

func (o IndexUsageExportTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndexUsageExportTaskInfo struct{}"
	}

	return strings.Join([]string{"IndexUsageExportTaskInfo", string(data)}, " ")
}
