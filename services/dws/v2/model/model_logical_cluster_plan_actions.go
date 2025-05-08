package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterPlanActions 逻辑集群增删计划行动
type LogicalClusterPlanActions struct {

	// 行动下一次触发时间
	NextFireTime *string `json:"next_fire_time,omitempty"`

	// 行动失败原因
	FailedReason *string `json:"failed_reason,omitempty"`

	// 行动ID
	Id *string `json:"id,omitempty"`

	// 行动类型，取值范围为(create|delete)
	Type *string `json:"type,omitempty"`

	// 行动周期性Cron表达式：如\"0 0 0 ? * 3\"
	Strategy *string `json:"strategy,omitempty"`

	// 行动状态，取值范围为(running|waiting|deleted|finished|disabled|failed)
	Status *string `json:"status,omitempty"`

	// 计划上一次触发时间
	PreFireTime *string `json:"pre_fire_time,omitempty"`
}

func (o LogicalClusterPlanActions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterPlanActions struct{}"
	}

	return strings.Join([]string{"LogicalClusterPlanActions", string(data)}, " ")
}
