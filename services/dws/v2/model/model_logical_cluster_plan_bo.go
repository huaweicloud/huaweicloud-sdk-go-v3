package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterPlanBo **参数解释**： 提交逻辑集群定时增删计划请求。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type LogicalClusterPlanBo struct {

	// **参数解释**： 逻辑集群名字。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// **参数解释**： 逻辑集群绑定的用户，若绑定了主逻辑集群，不能绑定用户。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	User *string `json:"user,omitempty"`

	// **参数解释**： 逻辑集群节点的个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NodeNum *int32 `json:"node_num,omitempty"`

	// **参数解释**： 逻辑集群的绑定的主逻辑集群，若绑定了用户，不能绑定主逻辑集群。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MainLogicalCluster *string `json:"main_logical_cluster,omitempty"`

	// **参数解释**： 计划类型，取值范围为(once|periodicity)。 **约束限制**： 不涉及。 **取值范围**： once：一次性计划 periodicity：周期性计划 **默认取值**： 不涉及。
	PlanType string `json:"plan_type"`

	// **参数解释**： 逻辑集群定时增删计划起始时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 逻辑集群定时增删计划终止时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 逻辑集群定时增删计划细节。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Actions []LogicalClusterPlanActionsParam `json:"actions"`
}

func (o LogicalClusterPlanBo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterPlanBo struct{}"
	}

	return strings.Join([]string{"LogicalClusterPlanBo", string(data)}, " ")
}
