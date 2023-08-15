package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListElasticResourcePoolQueuesRequest Request Object
type ListElasticResourcePoolQueuesRequest struct {

	// 弹性资源池名称
	ElasticResourcePoolName string `json:"elastic_resource_pool_name"`

	// 默认为100
	Limit *int32 `json:"limit,omitempty"`

	// 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 可以根据queueName进行过滤
	QueueName *string `json:"queue_name,omitempty"`
}

func (o ListElasticResourcePoolQueuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListElasticResourcePoolQueuesRequest struct{}"
	}

	return strings.Join([]string{"ListElasticResourcePoolQueuesRequest", string(data)}, " ")
}
