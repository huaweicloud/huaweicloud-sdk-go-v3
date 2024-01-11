package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadQueueUsersRequest Request Object
type ListWorkloadQueueUsersRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 队列名称
	QueueName string `json:"queue_name"`

	// 查询条数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListWorkloadQueueUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadQueueUsersRequest struct{}"
	}

	return strings.Join([]string{"ListWorkloadQueueUsersRequest", string(data)}, " ")
}
