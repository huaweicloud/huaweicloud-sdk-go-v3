package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowQueueRequest struct {
	// 待查询的队列ID

	QueueId string `json:"queue_id"`
	// 是否包含死信信息。  支持的值如下：  - true：包含死信消息。 - false：不包含死信消息。  默认值为：false。  Kafka队列没有死信功能，该参数对于Kafka队列无效。

	IncludeDeadletter *bool `json:"include_deadletter,omitempty"`
}

func (o ShowQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueueRequest struct{}"
	}

	return strings.Join([]string{"ShowQueueRequest", string(data)}, " ")
}
