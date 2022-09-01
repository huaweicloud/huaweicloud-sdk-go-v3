package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PutEventsRespEvents struct {

	// 发布失败的错误码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 发布失败的原因
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`

	// 事件唯一标识串
	EventId *string `json:"event_id,omitempty" xml:"event_id"`
}

func (o PutEventsRespEvents) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutEventsRespEvents struct{}"
	}

	return strings.Join([]string{"PutEventsRespEvents", string(data)}, " ")
}
