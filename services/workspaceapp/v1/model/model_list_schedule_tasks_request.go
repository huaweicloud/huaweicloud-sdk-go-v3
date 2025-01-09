package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduleTasksRequest Request Object
type ListScheduleTasksRequest struct {

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// 定时任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 任务类型： * `RESTART_SERVER` - 定时重启服务器 * `START_SERVER` - 定时开机 * `STOP_SERVER` - 定时关机 * `REINSTALL_OS` - 定时重装操作系统
	TaskType *string `json:"task_type,omitempty"`
}

func (o ListScheduleTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduleTasksRequest struct{}"
	}

	return strings.Join([]string{"ListScheduleTasksRequest", string(data)}, " ")
}
