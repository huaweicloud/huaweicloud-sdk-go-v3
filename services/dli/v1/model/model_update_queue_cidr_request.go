package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateQueueCidrRequest Request Object
type UpdateQueueCidrRequest struct {

	// 指定删除的队列名称。
	QueueName string `json:"queue_name"`

	Body *UpdateQueueCidrRequestBody `json:"body,omitempty"`
}

func (o UpdateQueueCidrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQueueCidrRequest struct{}"
	}

	return strings.Join([]string{"UpdateQueueCidrRequest", string(data)}, " ")
}
