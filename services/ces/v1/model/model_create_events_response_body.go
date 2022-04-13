package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 响应参数。
type CreateEventsResponseBody struct {
	// 事件ID。

	EventId string `json:"event_id"`
	// 事件名称。  必须以字母开头，只能包含0-9/a-z/A-Z/_，长度最短为1，最大为64。

	EventName string `json:"event_name"`
}

func (o CreateEventsResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventsResponseBody struct{}"
	}

	return strings.Join([]string{"CreateEventsResponseBody", string(data)}, " ")
}
