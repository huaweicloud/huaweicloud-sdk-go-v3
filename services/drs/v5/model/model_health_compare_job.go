package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HealthCompareJob 健康对比任务。
type HealthCompareJob struct {

	// 对比任务ID。
	Id *string `json:"id,omitempty"`

	// 对比类型： - object_comparison：对象对比。 - lines：行对比。 - account：用户对比。
	Type *string `json:"type,omitempty"`

	// 状态。 - WAITING_FOR_RUNNING：等待启动中 - RUNNING：运行中 - SUCCESSFUL：完成 - FAILED：失败 - CANCELLED：已取消 - TIMEOUT_INTERRUPT：超时中断 - FULL_DOING：全量校验中 - INCRE_DOING：增量校验中
	Status *string `json:"status,omitempty"`

	// 对比开始时间，UTC时间。
	StartTime *string `json:"start_time,omitempty"`

	// 对比结束时间，UTC时间。
	EndTime *string `json:"end_time,omitempty"`

	// 对比计算资源。
	ComputeType *string `json:"compute_type,omitempty"`
}

func (o HealthCompareJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthCompareJob struct{}"
	}

	return strings.Join([]string{"HealthCompareJob", string(data)}, " ")
}
