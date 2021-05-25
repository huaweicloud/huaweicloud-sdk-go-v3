package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListQueuesRequest struct {
	// 是否包含死信信息。  支持的值如下：  - true：包含死信消息。 - false：不包含死信消息。  默认值为：false。  Kafka队列没有死信功能，该参数对于Kafka队列无效。

	IncludeDeadletter *bool `json:"include_deadletter,omitempty"`
}

func (o ListQueuesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListQueuesRequest struct{}"
	}

	return strings.Join([]string{"ListQueuesRequest", string(data)}, " ")
}
