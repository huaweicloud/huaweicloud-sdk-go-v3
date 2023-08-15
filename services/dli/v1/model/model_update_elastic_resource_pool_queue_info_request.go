package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateElasticResourcePoolQueueInfoRequest Request Object
type UpdateElasticResourcePoolQueueInfoRequest struct {

	// 弹性资源池名称
	ElasticResourcePoolName string `json:"elastic_resource_pool_name"`

	// 队列名称
	QueueName string `json:"queue_name"`

	Body *UpdateElasticResourcePoolQueueScalingPolicyInfo `json:"body,omitempty"`
}

func (o UpdateElasticResourcePoolQueueInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateElasticResourcePoolQueueInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateElasticResourcePoolQueueInfoRequest", string(data)}, " ")
}
