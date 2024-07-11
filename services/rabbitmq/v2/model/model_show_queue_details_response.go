package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQueueDetailsResponse Response Object
type ShowQueueDetailsResponse struct {

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

	// 订阅该Queue的消费者信息。
	ConsumerDetails *[]ConsumerDetails `json:"consumer_details,omitempty"`

	// 以此Queue为目标的绑定信息列表。
	QueueBindings  *[]BindingsDetails `json:"queue_bindings,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowQueueDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueueDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowQueueDetailsResponse", string(data)}, " ")
}
