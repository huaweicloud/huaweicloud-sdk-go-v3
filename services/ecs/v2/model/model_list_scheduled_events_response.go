package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledEventsResponse Response Object
type ListScheduledEventsResponse struct {

	// 计划事件列表
	Events *[]EventResponse `json:"events,omitempty"`

	PageInfo       *ListEventsResponsePageInfo `json:"page_info,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListScheduledEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledEventsResponse struct{}"
	}

	return strings.Join([]string{"ListScheduledEventsResponse", string(data)}, " ")
}
