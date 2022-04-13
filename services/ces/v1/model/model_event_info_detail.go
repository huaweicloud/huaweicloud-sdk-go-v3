package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 一条事件监控详细信息
type EventInfoDetail struct {
	// 事件名称。  必须以字母开头，只能包含0-9/a-z/A-Z/_，长度最短为1，最大为64。

	EventName string `json:"event_name"`
	// 事件来源。 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，总长度最短为3，最大为32。

	EventSource *string `json:"event_source,omitempty"`
	// 事件发生时间。UNIX时间戳，单位毫秒。  说明： 因为客户端到服务器端有延时，因此插入数据的时间戳应该在[当前时间-1小时+20秒，当前时间+10分钟-20秒]区间内，保证到达服务器时不会因为传输时延造成数据不能插入数据库。

	Time int64 `json:"time"`

	Detail *EventItemDetail `json:"detail"`
	// 事件ID。

	EventId *string `json:"event_id,omitempty"`
}

func (o EventInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventInfoDetail struct{}"
	}

	return strings.Join([]string{"EventInfoDetail", string(data)}, " ")
}
