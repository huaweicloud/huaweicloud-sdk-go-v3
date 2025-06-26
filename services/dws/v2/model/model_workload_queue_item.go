package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadQueueItem **参数解释**： 资源池信息。 **取值范围**： 不涉及。
type WorkloadQueueItem struct {

	// **参数解释**： 资源池名称。 **取值范围**： 不涉及。
	QueueName string `json:"queue_name"`

	// **参数解释**： 逻辑集群名称。 **取值范围**： 不涉及。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// **参数解释**： 资源池短查询加速开关。 **取值范围**： 不涉及。
	ShortQueryOptimize *string `json:"short_query_optimize,omitempty"`

	// **参数解释**： 资源池短查询并发数。 **取值范围**： 不涉及。
	ShortQueryConcurrencyNum *int32 `json:"short_query_concurrency_num,omitempty"`

	// **参数解释**： 资源配置队列。 **取值范围**： 不涉及。
	ResourceItemList []WorkloadResourceItem `json:"resource_item_list"`
}

func (o WorkloadQueueItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadQueueItem struct{}"
	}

	return strings.Join([]string{"WorkloadQueueItem", string(data)}, " ")
}
