package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkloadQueueRequest Request Object
type ShowWorkloadQueueRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 资源队列名
	QueueName string `json:"queue_name"`
}

func (o ShowWorkloadQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkloadQueueRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkloadQueueRequest", string(data)}, " ")
}
