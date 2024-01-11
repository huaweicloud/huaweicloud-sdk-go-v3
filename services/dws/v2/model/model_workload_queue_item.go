package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadQueueItem 工作负载资源池
type WorkloadQueueItem struct {

	// 资源池名称。
	QueueName string `json:"queue_name"`

	// 逻辑集群名称。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 工作负载队列短查询加速开关。
	ShortQueryOptimize *string `json:"short_query_optimize,omitempty"`

	// 工作负载队列短查询并发数。
	ShortQueryConcurrencyNum *int32 `json:"short_query_concurrency_num,omitempty"`

	// 资源配置队列。
	ResourceItemList []WorkloadResourceItem `json:"resource_item_list"`
}

func (o WorkloadQueueItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadQueueItem struct{}"
	}

	return strings.Join([]string{"WorkloadQueueItem", string(data)}, " ")
}
