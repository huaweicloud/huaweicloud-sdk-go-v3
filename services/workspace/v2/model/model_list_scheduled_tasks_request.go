package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledTasksRequest Request Object
type ListScheduledTasksRequest struct {

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset int32 `json:"offset"`

	// 用于分页查询，每页返回的个数，取值范围0~50。
	Limit int32 `json:"limit"`

	// 任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 任务类型。START：开机，STOP：关机，REBOOT：重启，HIBERNATE：休眠，REBUILD：重建系统盘，EXECUTE_SCRIPT：执行脚本。
	TaskType *string `json:"task_type,omitempty"`

	// 执行周期类型。FIXED_TIME：指定时间，DAY：按天，WEEK：按周，MONTH：按月，LIFE_CYCLE：触发式。指定LIFE_CYCLE时，才返回触发式任务。
	ScheduledType *string `json:"scheduled_type,omitempty"`

	// 触发场景类型。POST_CREATE_DESKTOP_SUCCESS：创建桌面成功后，POST_REBUILD_DESKTOP_SUCCESS：重建桌面成功后，POST_REATTACH_DESKTOP_SUCCESS：触发重建的分配用户任务成功后。
	LifeCycleType *string `json:"life_cycle_type,omitempty"`

	// 最近一次执行状态。SUCCESS：成功，SKIP：跳过，FAIL：失败。
	LastStatus *string `json:"last_status,omitempty"`
}

func (o ListScheduledTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTasksRequest struct{}"
	}

	return strings.Join([]string{"ListScheduledTasksRequest", string(data)}, " ")
}
