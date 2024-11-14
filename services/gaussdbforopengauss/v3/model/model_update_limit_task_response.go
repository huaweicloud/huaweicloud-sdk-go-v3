package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLimitTaskResponse Response Object
type UpdateLimitTaskResponse struct {

	// 限流任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// CN节点数据库组,每个数据库字符串以逗号形式隔开，仅当任务类型为SQL_TYPE时，返回该值且与请求参数相同。
	Databases *string `json:"databases,omitempty"`

	// 限流任务名，与请求参数相同。
	TaskName *string `json:"task_name,omitempty"`

	// 关键词，仅当任务类型为SQL_TYPE时，返回该值且与请求参数相同。
	KeyWords *string `json:"key_words,omitempty"`

	// 并发数，与请求参数相同。
	ParallelSize *int32 `json:"parallel_size,omitempty"`

	// 限流任务开始时间，格式为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间，SQL范围返回该值。
	StartTime *string `json:"start_time,omitempty"`

	// 限流任务结束时间，格式为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间，SQL范围返回该值。
	EndTime *string `json:"end_time,omitempty"`

	// cpu利用率，仅当任务类型为SESSION_ACTIVE_MAX_COUNT时，返回该值且只保留正整数。
	CpuUtilization *int32 `json:"cpu_utilization,omitempty"`

	// 内存利用率，仅当任务类型为SESSION_ACTIVE_MAX_COUNT时，返回该值且只保留正整数。
	MemoryUtilization *int32 `json:"memory_utilization,omitempty"`

	// 规则名。
	RuleName *string `json:"rule_name,omitempty"`

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLimitTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLimitTaskResponse struct{}"
	}

	return strings.Join([]string{"UpdateLimitTaskResponse", string(data)}, " ")
}
