package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeScheduledEventRequest Request Object
type AuthorizeScheduledEventRequest struct {

	// 计划事件id。
	EventId string `json:"event_id"`

	Body *AuthorizeScheduledEventRequestBody `json:"body,omitempty"`
}

func (o AuthorizeScheduledEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeScheduledEventRequest struct{}"
	}

	return strings.Join([]string{"AuthorizeScheduledEventRequest", string(data)}, " ")
}
