package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLatestExportTaskByTypeResponse Response Object
type ShowLatestExportTaskByTypeResponse struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 导出任务状态，success:成功，failure:失败，running:运行中
	TaskStatus *string `json:"task_status,omitempty"`

	// 文件ID
	FileId *string `json:"file_id,omitempty"`

	// 文件名
	FileName       *string `json:"file_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowLatestExportTaskByTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLatestExportTaskByTypeResponse struct{}"
	}

	return strings.Join([]string{"ShowLatestExportTaskByTypeResponse", string(data)}, " ")
}
