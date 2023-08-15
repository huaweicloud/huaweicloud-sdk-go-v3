package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendImmediateEventResponse Response Object
type SendImmediateEventResponse struct {

	// **参数说明**：即时事件ID。
	EventId *string `json:"event_id,omitempty"`

	SendConfig *SendConfigResponse `json:"send_config,omitempty"`

	ImmediateEvent *ImmediateEventResponseDto `json:"immediate_event,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o SendImmediateEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendImmediateEventResponse struct{}"
	}

	return strings.Join([]string{"SendImmediateEventResponse", string(data)}, " ")
}
