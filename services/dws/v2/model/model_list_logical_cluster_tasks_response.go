package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicalClusterTasksResponse Response Object
type ListLogicalClusterTasksResponse struct {

	// **参数解释**： 逻辑集群任务信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogicalClusterTasks *[]LogicalClusterTaskInfo `json:"logical_cluster_tasks,omitempty"`

	// **参数解释**： 逻辑集群任务总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListLogicalClusterTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogicalClusterTasksResponse struct{}"
	}

	return strings.Join([]string{"ListLogicalClusterTasksResponse", string(data)}, " ")
}
