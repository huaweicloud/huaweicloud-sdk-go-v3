package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTask 导出任务
type ExportTask struct {

	// 导出任务id
	TaskId *string `json:"task_id,omitempty"`

	// 导出任务名称
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释：** 应用模板状态（域名粒度） **约束限制：** 不涉及 **取值范围：** - success: 应用模板成功 - fail: 应用模板失败 **默认取值：** 不涉及
	Status *string `json:"status,omitempty"`

	// 下载链接
	DownloadLink *string `json:"download_link,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 最近更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o ExportTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTask struct{}"
	}

	return strings.Join([]string{"ExportTask", string(data)}, " ")
}
