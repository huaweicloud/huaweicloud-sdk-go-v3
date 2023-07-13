package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlaybookInstanceMonitorDetail 剧本运行监控详情
type PlaybookInstanceMonitorDetail struct {

	// 运行总次数
	TotalInstanceRunNum *int32 `json:"total_instance_run_num,omitempty"`

	// 定时触发执行次数
	ScheduleInstanceRunNum *int32 `json:"schedule_instance_run_num,omitempty"`

	// 时间触发执行次数
	EventInstanceRunNum *int32 `json:"event_instance_run_num,omitempty"`

	MinRunTimeInstance *PlaybookInstanceRunStatistics `json:"min_run_time_instance,omitempty"`

	MaxRunTimeInstance *PlaybookInstanceRunStatistics `json:"max_run_time_instance,omitempty"`

	// 剧本实例总数
	TotalInstanceNum *int32 `json:"total_instance_num,omitempty"`

	// 运行成功实例数量
	SuccessInstanceNum *int32 `json:"success_instance_num,omitempty"`

	// 运行失败实例数量
	FailInstanceNum *int32 `json:"fail_instance_num,omitempty"`

	// 运行终止实例数量
	TerminateInstanceNum *int32 `json:"terminate_instance_num,omitempty"`

	// 运行中实例数量
	RunningInstanceNum *int32 `json:"running_instance_num,omitempty"`
}

func (o PlaybookInstanceMonitorDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlaybookInstanceMonitorDetail struct{}"
	}

	return strings.Join([]string{"PlaybookInstanceMonitorDetail", string(data)}, " ")
}
