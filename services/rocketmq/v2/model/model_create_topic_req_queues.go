package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTopicReqQueues struct {

	// 关联的代理。
	Broker *string `json:"broker,omitempty"`

	// 队列数，范围1~50。
	QueueNum float32 `json:"queue_num,omitempty"`
}

func (o CreateTopicReqQueues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicReqQueues struct{}"
	}

	return strings.Join([]string{"CreateTopicReqQueues", string(data)}, " ")
}
