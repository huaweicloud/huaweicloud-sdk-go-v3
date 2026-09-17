package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIndexUsageExportTasksResponse Response Object
type ListIndexUsageExportTasksResponse struct {

	// 总数
	Total *int64 `json:"total,omitempty"`

	// 任务列表
	TaskList       *[]IndexUsageExportTaskInfo `json:"task_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListIndexUsageExportTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIndexUsageExportTasksResponse struct{}"
	}

	return strings.Join([]string{"ListIndexUsageExportTasksResponse", string(data)}, " ")
}
