package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueryExportTaskResponse Response Object
type ListQueryExportTaskResponse struct {

	// **参数解释**： 任务ID **取值范围**： 字符长度1-64位
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**： 任务名称 **取值范围**： 字符长度1-128位
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**： 导出任务状态 **取值范围**： - success：成功 - failure：失败 - running：运行中
	TaskStatus *string `json:"task_status,omitempty"`

	// **参数解释**： 文件ID **取值范围**： 字符长度1-64位
	FileId *string `json:"file_id,omitempty"`

	// **参数解释**： 文件名 **取值范围**： 字符长度1-255位
	FileName       *string `json:"file_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListQueryExportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryExportTaskResponse struct{}"
	}

	return strings.Join([]string{"ListQueryExportTaskResponse", string(data)}, " ")
}
