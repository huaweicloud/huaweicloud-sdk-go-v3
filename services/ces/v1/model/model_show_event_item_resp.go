package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEventItemResp 请求参数。
type ShowEventItemResp struct {

	// **参数解释**： 事件名称。 **取值范围**： 必须以字母开头，只能包含0-9/a-z/A-Z/_，长度为[1,64]个字符。
	EventName string `json:"event_name"`

	// **参数解释**： 事件来源。 **取值范围**： 格式为service.item。service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_，长度为[3,32]个字符。
	EventSource string `json:"event_source"`

	// **参数解释**： 事件发生时间。UNIX时间戳，单位毫秒。 **取值范围**： 因为客户端到服务器端有延时，因此插入数据的时间戳应该在[当前时间-1小时+20秒,当前时间+10分钟-20秒]区间内，保证到达服务器时不会因为传输时延造成数据不能插入数据库。
	Time int64 `json:"time"`

	Detail *ShowEventItemDetailResp `json:"detail"`
}

func (o ShowEventItemResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventItemResp struct{}"
	}

	return strings.Join([]string{"ShowEventItemResp", string(data)}, " ")
}
