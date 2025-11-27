package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExportTasksRequest Request Object
type ListExportTasksRequest struct {

	// - 每页显示的条目数量, 默认为10
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 任务id
	TaskId string `json:"task_id"`

	// 任务名称
	TaskName string `json:"task_name"`
}

func (o ListExportTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExportTasksRequest struct{}"
	}

	return strings.Join([]string{"ListExportTasksRequest", string(data)}, " ")
}
