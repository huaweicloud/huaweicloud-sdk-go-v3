package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTask 导出任务
type ExportTask struct {

	// **参数解释：** 导出任务id **取值范围：** 不涉及
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释：** 导出任务名称 **取值范围：** 不涉及
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释：** 任务状态 **约束限制：** 不涉及 **取值范围：** - success: 成功 - fail: 失败 **默认取值：** 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释：** 下载链接 **取值范围：** 不涉及
	DownloadLink *string `json:"download_link,omitempty"`

	// **参数解释：** 创建时间 **取值范围：** 不涉及
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释：** 最近更新时间 **取值范围：** 不涉及
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o ExportTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTask struct{}"
	}

	return strings.Join([]string{"ExportTask", string(data)}, " ")
}
