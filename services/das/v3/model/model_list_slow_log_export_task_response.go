package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSlowLogExportTaskResponse Response Object
type ListSlowLogExportTaskResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 导出任务列表
	TaskList       *[]SlowLogExportTask `json:"task_list,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListSlowLogExportTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogExportTaskResponse struct{}"
	}

	return strings.Join([]string{"ListSlowLogExportTaskResponse", string(data)}, " ")
}
