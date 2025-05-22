package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadQueue **参数解释**： 资源池信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type WorkloadQueue struct {

	// **参数解释**： 工作负载队列名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	WorkloadQueueName string `json:"workload_queue_name"`

	// **参数解释**： 逻辑集群名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// **参数解释**： 资源配置队列。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	WorkloadResourceItemList []WorkloadResource `json:"workload_resource_item_list"`
}

func (o WorkloadQueue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadQueue struct{}"
	}

	return strings.Join([]string{"WorkloadQueue", string(data)}, " ")
}
