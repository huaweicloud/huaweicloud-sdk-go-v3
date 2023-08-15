package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IefEvents IEF系统事件的数据
type IefEvents struct {

	// 资源类型
	EventType string `json:"event_type"`

	// 资源的操作类型
	Operation string `json:"operation"`

	// 事件产生的时间戳
	Timestamp int64 `json:"timestamp"`

	// 消息发送的Topic
	Topic string `json:"topic"`

	// 资源名称
	Name string `json:"name"`

	// 资源的属性
	Attributes *string `json:"attributes,omitempty"`
}

func (o IefEvents) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IefEvents struct{}"
	}

	return strings.Join([]string{"IefEvents", string(data)}, " ")
}
