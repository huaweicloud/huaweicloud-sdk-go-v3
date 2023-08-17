package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RtcCause struct {

	// 异常事件的时间戳
	Ts *string `json:"ts,omitempty"`

	// 异常事件ID
	EventId *string `json:"event_id,omitempty"`

	// 对端的用户ID
	PeerId *string `json:"peer_id,omitempty"`
}

func (o RtcCause) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RtcCause struct{}"
	}

	return strings.Join([]string{"RtcCause", string(data)}, " ")
}
