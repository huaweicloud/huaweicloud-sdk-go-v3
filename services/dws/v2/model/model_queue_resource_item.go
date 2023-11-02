package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueueResourceItem 工作负载资源池
type QueueResourceItem struct {

	// 资源池名称。
	QueueName string `json:"queue_name"`

	// 资源配置队列。
	QueueResources []WorkloadResourceItem `json:"queue_resources"`
}

func (o QueueResourceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueueResourceItem struct{}"
	}

	return strings.Join([]string{"QueueResourceItem", string(data)}, " ")
}
