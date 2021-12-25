package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListMonitorLogRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 任务ID

	TaskId string `json:"task_id"`
	// 偏移量，表示从此偏移量开始查询， offset大于等于1

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示条目数量，最大数量999，超过999后只返回999

	Limit *int32 `json:"limit,omitempty"`
	// 日志查询的起始时间，格式timestamp(ms)，使用UTC时区

	BeginTime *int64 `json:"begin_time,omitempty"`
	// 日志查询的结束时间，格式timestamp(ms)，使用UTC时区

	EndTime *int64 `json:"end_time,omitempty"`
}

func (o ListMonitorLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitorLogRequest struct{}"
	}

	return strings.Join([]string{"ListMonitorLogRequest", string(data)}, " ")
}
