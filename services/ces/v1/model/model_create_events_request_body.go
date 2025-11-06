package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEventsRequestBody 请求参数。
type CreateEventsRequestBody struct {

	// **参数解释**： 事件名称。 **约束限制**： 不涉及。 **取值范围**： 必须以字母开头，只能包含0-9/a-z/A-Z/_，长度为[1,64]个字符。 **默认取值**： 不涉及。
	EventName string `json:"event_name"`

	// **参数解释**： 事件来源。 **约束限制**： 不涉及。 **取值范围**： 格式为service.item，根据实际情况自定义配置。 service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，service.item，长度为[3,32]个字符。 **默认取值**： 不涉及。
	EventSource string `json:"event_source"`

	// **参数解释**： 事件发生时间。UNIX时间戳，单位毫秒。 **约束限制**： 因为客户端到服务器端有延时，因此插入数据的时间戳应该在[当前时间-1小时+10秒，当前时间+10分钟-10秒]区间内，保证到达服务器时不会因为传输时延造成数据不能插入数据库。 **取值范围**： 插入数据的时间戳应该在[当前时间-1小时+10秒，当前时间+10分钟-10秒]区间内。 **默认取值**： 不涉及。
	Time int64 `json:"time"`

	Detail *EventItemDetail `json:"detail"`
}

func (o CreateEventsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEventsRequestBody", string(data)}, " ")
}
