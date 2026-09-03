package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFullSqlExportTasksResponse Response Object
type ListFullSqlExportTasksResponse struct {

	// 导出任务总数
	Total *int32 `json:"total,omitempty"`

	// 导出任务列表
	TaskList       *[]FullSqlExportTaskInfo `json:"task_list,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListFullSqlExportTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFullSqlExportTasksResponse struct{}"
	}

	return strings.Join([]string{"ListFullSqlExportTasksResponse", string(data)}, " ")
}
