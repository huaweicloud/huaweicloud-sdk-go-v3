package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueuePropertyRequest Request Object
type ListQueuePropertyRequest struct {

	// 队列名称
	QueueName string `json:"queue_name"`

	// 列表当前页
	Page *int32 `json:"page,omitempty"`

	// 每页显示条数
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o ListQueuePropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueuePropertyRequest struct{}"
	}

	return strings.Join([]string{"ListQueuePropertyRequest", string(data)}, " ")
}
