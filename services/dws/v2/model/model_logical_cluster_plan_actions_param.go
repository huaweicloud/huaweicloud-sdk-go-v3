package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterPlanActionsParam 用于提交逻辑集群增删计划的行为信息
type LogicalClusterPlanActionsParam struct {

	// 定时增删计划行为类型，取值范围为（create|delete）
	Type string `json:"type"`

	// 周期性定时增删计划，Cron策略表达式：如\"0 0 0 ? * 3\"
	Strategy *string `json:"strategy,omitempty"`
}

func (o LogicalClusterPlanActionsParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterPlanActionsParam struct{}"
	}

	return strings.Join([]string{"LogicalClusterPlanActionsParam", string(data)}, " ")
}
