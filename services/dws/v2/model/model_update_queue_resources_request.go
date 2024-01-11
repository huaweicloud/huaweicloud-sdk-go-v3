package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateQueueResourcesRequest Request Object
type UpdateQueueResourcesRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 队列名称
	QueueName string `json:"queue_name"`

	Body *WorkloadQueueRequest `json:"body,omitempty"`
}

func (o UpdateQueueResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQueueResourcesRequest struct{}"
	}

	return strings.Join([]string{"UpdateQueueResourcesRequest", string(data)}, " ")
}
