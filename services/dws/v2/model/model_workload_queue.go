package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadQueue 工作负载队列
type WorkloadQueue struct {

	// 工作负载队列名称。
	WorkloadQueueName string `json:"workload_queue_name"`

	// 逻辑集群名称。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 资源配置队列。
	WorkloadResourceItemList []WorkloadResource `json:"workload_resource_item_list"`
}

func (o WorkloadQueue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadQueue struct{}"
	}

	return strings.Join([]string{"WorkloadQueue", string(data)}, " ")
}
