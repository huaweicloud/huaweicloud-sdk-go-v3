package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ConsumeMessagesRequest struct {
	// 指定的队列ID。

	QueueId string `json:"queue_id"`
	// 消费组的ID。

	ConsumerGroupId string `json:"consumer_group_id"`
	// 获取可消费的消息的条数。  取值范围：1~10。  默认值：10

	MaxMsgs *int32 `json:"max_msgs,omitempty"`
	// 设定队列可消费的消息为0时的读取消息等待时间。  如果在等待时间内有新的消息，则立即返回消费结果，如果等待时间内没有新的消息，则到等待时间后返回消费结果。  取值范围：1~60s  默认值：3s  说明：不带该参数或者配置为空，都默认为3s。

	TimeWait *int32 `json:"time_wait,omitempty"`
	// 提交确认消费的超时时间，客户端需要在该时间内提交消费确认，如果超过指定时间，没有确认消费，系统会报消息确认超时或handler无效，则默认为消费失败。  取值范围：15~300s  默认值：30s  说明：不带该参数或者配置为空，都默认为30s。

	AckWait *int32 `json:"ack_wait,omitempty"`
	// 添加标签后可以按照Tag进行过滤，只消费匹配上标签的消息。  Tag的数量不超过3个。  每个Tag长度不超过64。

	Tag *string `json:"tag,omitempty"`
	// 多个消息标签的过滤类型。  取值范围： - and：必须所有标签匹配上，才能消费消息。 - or：只要有一条标签匹配上，就可以消费消息。  默认值为：or。

	TagType *string `json:"tag_type,omitempty"`
}

func (o ConsumeMessagesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ConsumeMessagesRequest struct{}"
	}

	return strings.Join([]string{"ConsumeMessagesRequest", string(data)}, " ")
}
