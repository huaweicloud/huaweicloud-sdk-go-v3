package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PutEventsRespEvents struct {

	// 发布失败的错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 发布失败的原因
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 事件唯一标识串
	EventId *string `json:"event_id,omitempty"`
}

func (o PutEventsRespEvents) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutEventsRespEvents struct{}"
	}

	return strings.Join([]string{"PutEventsRespEvents", string(data)}, " ")
}
