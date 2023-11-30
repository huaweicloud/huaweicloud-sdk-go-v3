package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueuePropertiesRequest Request Object
type ListQueuePropertiesRequest struct {

	// 队列名称
	QueueName string `json:"queue_name"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示条数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListQueuePropertiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueuePropertiesRequest struct{}"
	}

	return strings.Join([]string{"ListQueuePropertiesRequest", string(data)}, " ")
}
