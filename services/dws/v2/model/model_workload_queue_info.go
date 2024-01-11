package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadQueueInfo 资源池
type WorkloadQueueInfo struct {

	// 资源池名称。
	WorkloadQueueName string `json:"workload_queue_name"`

	// 逻辑集群名称。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 资源配置队列。
	ResourceItemList []WorkloadResourceItem `json:"resource_item_list"`
}

func (o WorkloadQueueInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadQueueInfo struct{}"
	}

	return strings.Join([]string{"WorkloadQueueInfo", string(data)}, " ")
}
