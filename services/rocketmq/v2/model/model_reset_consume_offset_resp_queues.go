package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResetConsumeOffsetRespQueues struct {

	// 队列所在的broker。
	BrokerName *string `json:"broker_name,omitempty"`

	// 队列ID。
	QueueId *int32 `json:"queue_id,omitempty"`

	// 重置消费进度。
	TimestampOffset *int64 `json:"timestamp_offset,omitempty"`
}

func (o ResetConsumeOffsetRespQueues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConsumeOffsetRespQueues struct{}"
	}

	return strings.Join([]string{"ResetConsumeOffsetRespQueues", string(data)}, " ")
}
