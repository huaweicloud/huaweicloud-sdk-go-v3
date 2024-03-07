package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkloadQueueRequest Request Object
type DeleteWorkloadQueueRequest struct {

	// 集群ID。
	ClusterId string `json:"cluster_id"`

	// 逻辑集群名称。逻辑集群模式下该字段必填。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 工作负载队列名称。
	WorkloadQueueName string `json:"workload_queue_name"`
}

func (o DeleteWorkloadQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkloadQueueRequest struct{}"
	}

	return strings.Join([]string{"DeleteWorkloadQueueRequest", string(data)}, " ")
}
