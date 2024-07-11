package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueueDetails struct {

	// Queue所属Vhost名称
	Vhost *string `json:"vhost,omitempty"`

	// Queue名称
	Name *string `json:"name,omitempty"`

	// 是否持久化
	Durable *bool `json:"durable,omitempty"`

	// 是否自动删除
	AutoDelete *bool `json:"auto_delete,omitempty"`

	// 待消费消息数
	Messages *int32 `json:"messages,omitempty"`

	// 连接的消费者数
	Consumers *int32 `json:"consumers,omitempty"`

	// 策略[（AMQP版本不支持policy，不涉及此参数）](tag:hws,hws_hk)
	Policy *string `json:"policy,omitempty"`

	Arguments *QueueArguments `json:"arguments,omitempty"`
}

func (o QueueDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueueDetails struct{}"
	}

	return strings.Join([]string{"QueueDetails", string(data)}, " ")
}
