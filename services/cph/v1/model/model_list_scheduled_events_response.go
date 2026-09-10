package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledEventsResponse Response Object
type ListScheduledEventsResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 计划事件总数。
	Count *int32 `json:"count,omitempty"`

	// 计划事件信息
	ScheduledEvents *[]ListScheduledEventsResponseBodyScheduledEvents `json:"scheduled_events,omitempty"`

	PageInfo       *ListCloudPhoneServersModelOfferingsResponseBodyPageInfo `json:"page_info,omitempty"`
	HttpStatusCode int                                                      `json:"-"`
}

func (o ListScheduledEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledEventsResponse struct{}"
	}

	return strings.Join([]string{"ListScheduledEventsResponse", string(data)}, " ")
}
