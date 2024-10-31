package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLimitTaskResponse Response Object
type ShowLimitTaskResponse struct {

	// 限流任务名传。
	TaskName *string `json:"task_name,omitempty"`

	// 任务限流类型。
	LimitType *string `json:"limit_type,omitempty"`

	// 并发数。
	ParallelSize *int32 `json:"parallel_size,omitempty"`

	// 限流任务开始时间,格式为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间。
	StartTime *string `json:"start_time,omitempty"`

	// 限流任务结束时间,格式为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间。
	EndTime *string `json:"end_time,omitempty"`

	// 限流任务已执行时间，单位秒。
	TaskRunningTime *int32 `json:"task_running_time,omitempty"`

	// 限流任务拦截次数。
	LimitCount *int32 `json:"limit_count,omitempty"`

	// 规则名。
	RuleName *string `json:"rule_name,omitempty"`

	// 内存利用率，仅当任务类型为SESSION_ACTIVE_MAX_COUNT时，返回该值且与请求参数相同。
	MemoryUtilization *int32 `json:"memory_utilization,omitempty"`

	// cpu利用率，仅当任务类型为SESSION_ACTIVE_MAX_COUNT时，返回该值且与请求参数相同。
	CpuUtilization *int32 `json:"cpu_utilization,omitempty"`

	// 限流任务列表
	LimitTaskRuleInfoList *[]LimitTaskRuleInfoOption `json:"limit_task_rule_info_list,omitempty"`
	HttpStatusCode        int                        `json:"-"`
}

func (o ShowLimitTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLimitTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowLimitTaskResponse", string(data)}, " ")
}
