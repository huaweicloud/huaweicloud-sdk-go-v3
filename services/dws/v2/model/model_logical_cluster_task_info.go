package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterTaskInfo **参数解释**： 逻辑集群任务信息。 **取值范围**： 不涉及。
type LogicalClusterTaskInfo struct {

	// **参数解释**： 任务类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 逻辑集群名称。 **取值范围**： 不涉及。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// **参数解释**： 任务开始时间。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 任务结束时间。 **取值范围**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 任务执行结果。 **取值范围**： 不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释**： 任务执行日志。 **取值范围**： 不涉及。
	Log *string `json:"log,omitempty"`
}

func (o LogicalClusterTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterTaskInfo struct{}"
	}

	return strings.Join([]string{"LogicalClusterTaskInfo", string(data)}, " ")
}
