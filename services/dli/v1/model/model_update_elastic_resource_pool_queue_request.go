package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateElasticResourcePoolQueueRequest Request Object
type UpdateElasticResourcePoolQueueRequest struct {

	// 弹性资源池名称
	ElasticResourcePoolName string `json:"elastic_resource_pool_name"`

	// 队列名称
	QueueName string `json:"queue_name"`

	Body *UpdateElasticResourcePoolQueueRequestBody `json:"body,omitempty"`
}

func (o UpdateElasticResourcePoolQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateElasticResourcePoolQueueRequest struct{}"
	}

	return strings.Join([]string{"UpdateElasticResourcePoolQueueRequest", string(data)}, " ")
}
