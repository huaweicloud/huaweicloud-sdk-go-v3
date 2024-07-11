package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueueArguments Queue参数，如果未配置则不返回。
type QueueArguments struct {

	// 消息过期时间，发布到Queue的消息在被丢弃之前可以存活多长时间。
	XMessageTtl *int64 `json:"x-message-ttl,omitempty"`

	// 死信Exchange名称，消息被拒绝或过期时将重新发布到该Exchange。
	XDeadLetterExchange *string `json:"x-dead-letter-exchange,omitempty"`

	// 死信的RoutingKey，死信Exchange会发送死信消息到绑定对应RoutingKey的Queue上。
	XDeadLetterRoutingKey *string `json:"x-dead-letter-routing-key,omitempty"`

	// 惰性队列[（AMQP版本默认持久化所有消息，不涉及此参数）](tag:hws,hws_hk)
	XQueueMode *string `json:"x-queue-mode,omitempty"`
}

func (o QueueArguments) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueueArguments struct{}"
	}

	return strings.Join([]string{"QueueArguments", string(data)}, " ")
}
