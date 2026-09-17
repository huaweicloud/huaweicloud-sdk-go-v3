package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMissingIndexExportTasksResponse Response Object
type ListMissingIndexExportTasksResponse struct {

	// 总数
	Total *int64 `json:"total,omitempty"`

	// 任务列表
	TaskList       *[]MissingIndexExportTaskInfo `json:"task_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListMissingIndexExportTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMissingIndexExportTasksResponse struct{}"
	}

	return strings.Join([]string{"ListMissingIndexExportTasksResponse", string(data)}, " ")
}
