package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEventSubRequest Request Object
type UpdateEventSubRequest struct {

	// 事件订阅ID
	EventSubId string `json:"event_sub_id"`

	Body *EventSubUpdateRequest `json:"body,omitempty"`
}

func (o UpdateEventSubRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEventSubRequest struct{}"
	}

	return strings.Join([]string{"UpdateEventSubRequest", string(data)}, " ")
}
