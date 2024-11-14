package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLimitTaskResponse Response Object
type CreateLimitTaskResponse struct {

	// 限流任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 任务限流范围，与请求参数相同。
	TaskScope *string `json:"task_scope,omitempty"`

	// 任务限流类型，与请求参数相同。
	LimitType *string `json:"limit_type,omitempty"`

	// 任务限流类型值，与请求参数相同。
	LimitTypeValue *string `json:"limit_type_value,omitempty"`

	// CN节点数据库组，每个数据库字符串以逗号形式隔开。
	Databases *string `json:"databases,omitempty"`

	// 限流任务名。
	TaskName *string `json:"task_name,omitempty"`

	// SQL模板,仅当任务类型为SQL_ID时，返回该值且与请求参数相同。
	SqlModel *string `json:"sql_model,omitempty"`

	// 关键词，仅当任务类型为SQL_TYPE时，返回该值且与请求参数相同。
	KeyWords *string `json:"key_words,omitempty"`

	// 限流任务状态，当前支持：CREATING，UPDATEING，DELETING，WAIT_EXCUTE，EXCUTING，TIME_OVER，DELETED，CREATE_FAILED，UPDATE_FAILED，DELETE_FAILED，EXCEPTION，NODE_SHUT_DOWN。
	Status *string `json:"status,omitempty"`

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 规则名。
	RuleName *string `json:"rule_name,omitempty"`

	// 并发数，与请求参数相同。
	ParallelSize *int32 `json:"parallel_size,omitempty"`

	// cpu利用率阈值，仅当任务类型为SESSION_ACTIVE_MAX_COUNT时，返回该值且只保留正整数。
	CpuUtilization *int32 `json:"cpu_utilization,omitempty"`

	// 内存利用率阈值，仅当任务类型为SESSION_ACTIVE_MAX_COUNT时，返回该值且只保留正整数。
	MemoryUtilization *int32 `json:"memory_utilization,omitempty"`

	// 限流任务开始时间，与请求参数相同，格式为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间。
	StartTime *string `json:"start_time,omitempty"`

	// 限流任务结束时间，与请求参数相同，格式为yyyy-mm-ddThh:mm:ssZ，当前时间指UTC时间。
	EndTime *string `json:"end_time,omitempty"`

	// 创建时间为本地时间，格式为yyyy-mm-ddThh:mm:ssZ。
	Created *string `json:"created,omitempty"`

	// 更新时间为本地时间，格式为yyyy-mm-ddThh:mm:ssZ。
	Updated *string `json:"updated,omitempty"`

	// 创建者。
	Creator *string `json:"creator,omitempty"`

	// 更新者。
	Modifier *string `json:"modifier,omitempty"`

	// CN节点信息列表，如果类型为SQLID，必返。
	NodeInfos *[]CreateLimitTaskNodeResult `json:"node_infos,omitempty"`

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLimitTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLimitTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateLimitTaskResponse", string(data)}, " ")
}
