package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogicalClusterPlanActions 更新逻辑集群增删计划细节信息
type UpdateLogicalClusterPlanActions struct {

	// 更新逻辑集群增删行动ID
	Id *string `json:"id,omitempty"`

	// 更新逻辑集群增删计划行动类型，取值范围为(create|delete)
	Type *string `json:"type,omitempty"`

	// 更新逻辑集群增删计划行为Cron策略表达式
	Strategy *string `json:"strategy,omitempty"`
}

func (o UpdateLogicalClusterPlanActions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogicalClusterPlanActions struct{}"
	}

	return strings.Join([]string{"UpdateLogicalClusterPlanActions", string(data)}, " ")
}
