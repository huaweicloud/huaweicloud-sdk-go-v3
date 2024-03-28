package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListElasticResourcePoolQueuesResponse Response Object
type ListElasticResourcePoolQueuesResponse struct {

	// 是否成功
	IsSuccess *bool `json:"is_success,omitempty"`

	// 请求消息
	Message *string `json:"message,omitempty"`

	// 该弹性资源池下所有queue信息及队列扩缩容策略信息。
	Queues *[]ElasticResourcePoolQueue `json:"queues,omitempty"`

	// 该资源池下关联的队列数量
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListElasticResourcePoolQueuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListElasticResourcePoolQueuesResponse struct{}"
	}

	return strings.Join([]string{"ListElasticResourcePoolQueuesResponse", string(data)}, " ")
}
