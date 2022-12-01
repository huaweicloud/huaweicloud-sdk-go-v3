package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 工作负载队列
type WorkloadQueue struct {

	// 工作负载队列名称。
	WorkloadQueueName *string `json:"workload_queue_name,omitempty"`

	// 逻辑集群名称。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 短查询开关。
	ShortQueryOptimize *string `json:"short_query_optimize,omitempty"`

	// 短查询并发数。
	ShortQueryConcurrencyNum *string `json:"short_query_concurrency_num,omitempty"`

	// 资源配置队列。
	WorkloadResourceItemList *[]WorkloadResource `json:"workload_resource_item_list,omitempty"`
}

func (o WorkloadQueue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadQueue struct{}"
	}

	return strings.Join([]string{"WorkloadQueue", string(data)}, " ")
}
