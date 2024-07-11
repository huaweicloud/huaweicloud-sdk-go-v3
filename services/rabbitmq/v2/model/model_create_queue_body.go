package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateQueueBody struct {

	// Queue名称
	Name string `json:"name"`

	// 是否自动删除
	AutoDelete bool `json:"auto_delete"`

	// 是否持久化[（AMQP版本默认持久化，不涉及此字段）](tag:hws,hws_hk)
	Durable *bool `json:"durable,omitempty"`

	// 死信Exchange名称，消息被拒绝或过期时将重新发布到该Exchange。
	DeadLetterExchange *string `json:"dead_letter_exchange,omitempty"`

	// 死信Exchange的RoutingKey，死信Exchange会发送死信消息到绑定对应RoutingKey的Queue上。
	DeadLetterRoutingKey *string `json:"dead_letter_routing_key,omitempty"`

	// 发布到Queue的消息在被丢弃之前可以存活多长时间
	MessageTtl *int64 `json:"message_ttl,omitempty"`

	// 若设置惰性队列，请输入lazy。惰性队列模式会在磁盘上存储尽可能多的消息以减少内存使用；若不设置，队列将消息存储在内存缓存以尽可能快地传递消息。[（AMQP版本默认将消息存储到磁盘，不涉及此字段）](tag:hws,hws_hk)
	LazyMode *string `json:"lazy_mode,omitempty"`
}

func (o CreateQueueBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueueBody struct{}"
	}

	return strings.Join([]string{"CreateQueueBody", string(data)}, " ")
}
