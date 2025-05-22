package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterPlanVo **参数解释**： 逻辑集群增删计划对象。 **取值范围**： 不涉及。
type LogicalClusterPlanVo struct {

	// **参数解释**： 增删逻辑集群计划ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 逻辑集群名字。 **取值范围**： 不涉及。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// **参数解释**： 逻辑集群节点个数。 **取值范围**： 不涉及。
	NodeNum *int32 `json:"node_num,omitempty"`

	// **参数解释**： 逻辑集群增删计划类型，取值范围为 (once|periodicity)，表示单次或周期性。 **取值范围**： 不涉及。
	PlanType *string `json:"plan_type,omitempty"`

	// **参数解释**： 逻辑集群增删计划状态。 **取值范围**： running：运行中 waiting：等待中 deleted：已删除 finished：已完成 disabled：已禁用 failed：失败
	Status *string `json:"status,omitempty"`

	// **参数解释**： 逻辑集群增删计划开始时间。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 逻辑集群增删计划结束时间。 **取值范围**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 逻辑集群增删计划更新时间。 **取值范围**： 不涉及。
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**： 逻辑集群增删计划绑定的用户。 **取值范围**： 不涉及。
	User *string `json:"user,omitempty"`

	// **参数解释**： 任务信息。 **取值范围**： 不涉及。
	Actions *[]LogicalClusterPlanActions `json:"actions,omitempty"`
}

func (o LogicalClusterPlanVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterPlanVo struct{}"
	}

	return strings.Join([]string{"LogicalClusterPlanVo", string(data)}, " ")
}
