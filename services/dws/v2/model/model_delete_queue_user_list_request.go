package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteQueueUserListRequest Request Object
type DeleteQueueUserListRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 队列名称
	QueueName string `json:"queue_name"`

	Body *WorkloadQueueUserReq `json:"body,omitempty"`
}

func (o DeleteQueueUserListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueueUserListRequest struct{}"
	}

	return strings.Join([]string{"DeleteQueueUserListRequest", string(data)}, " ")
}
