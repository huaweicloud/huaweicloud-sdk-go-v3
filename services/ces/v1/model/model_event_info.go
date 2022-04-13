package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 一条事件监控信息
type EventInfo struct {
	// 事件名称。

	EventName *string `json:"event_name,omitempty"`
	// 事件类型。

	EventType *string `json:"event_type,omitempty"`
	// 选择查询的时间范围内，此事件发生的数量。

	EventCount *int32 `json:"event_count,omitempty"`
	// 此事件最近一次发生的时间。

	LatestOccurTime *int64 `json:"latest_occur_time,omitempty"`
	// 事件来源，如果是系统事件则值为各服务的命名空间，各服务的命名空间可查看：“[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”；如果是自定义事件，则为用户自定义上报定义。

	LatestEventSource *string `json:"latest_event_source,omitempty"`
}

func (o EventInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventInfo struct{}"
	}

	return strings.Join([]string{"EventInfo", string(data)}, " ")
}
