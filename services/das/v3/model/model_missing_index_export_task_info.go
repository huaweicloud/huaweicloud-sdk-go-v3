package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MissingIndexExportTaskInfo 缺失索引导出任务信息
type MissingIndexExportTaskInfo struct {

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

func (o MissingIndexExportTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MissingIndexExportTaskInfo struct{}"
	}

	return strings.Join([]string{"MissingIndexExportTaskInfo", string(data)}, " ")
}
