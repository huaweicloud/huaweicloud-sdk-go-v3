package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowQueueRequest struct {
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`
	// 队列ID，用于唯一标识一个队列。

	QueueId string `json:"queue_id"`
}

func (o ShowQueueRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQueueRequest struct{}"
	}

	return strings.Join([]string{"ShowQueueRequest", string(data)}, " ")
}
