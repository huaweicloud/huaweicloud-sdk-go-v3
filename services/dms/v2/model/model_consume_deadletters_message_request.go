package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ConsumeDeadlettersMessageRequest struct {
	// 指定的队列ID。

	QueueId string `json:"queue_id"`
	// 消费组的ID。

	ConsumerGroupId string `json:"consumer_group_id"`
	// 获取可消费的死信消息的条数。  >单次消费返回的消息数量可能会少于指定条数，但多次消费最终可获取全部消息。  取值范围：1~10。  默认值：10

	MaxMsgs *int32 `json:"max_msgs,omitempty"`
	// 设定消费组中可消费的死信为0时的读取消息等待时间。  如果在等待时间内有新的死信消息，则立即返回消费结果，如果等待时间内没有新的死信消息，则到等待时间后返回消费结果。  取值范围：1~60s  默认值：3s  说明：不带该参数或者配置为空，都默认为3s。

	TimeWait *int32 `json:"time_wait,omitempty"`
	// commit提交超时时间，在该时间内提交确认，确认有效，如果超过指定时间，系统会报消息确认超时，或handler无效。  取值范围：15~300s  默认值：30s  说明：不带该参数或者配置为空，都默认为30s。

	AckWait *int32 `json:"ack_wait,omitempty"`
}

func (o ConsumeDeadlettersMessageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ConsumeDeadlettersMessageRequest struct{}"
	}

	return strings.Join([]string{"ConsumeDeadlettersMessageRequest", string(data)}, " ")
}
