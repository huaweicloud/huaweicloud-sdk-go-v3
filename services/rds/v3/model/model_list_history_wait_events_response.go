package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryWaitEventsResponse Response Object
type ListHistoryWaitEventsResponse struct {

	// 总记录数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 事件列表
	Events         *[]Event `json:"events,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListHistoryWaitEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryWaitEventsResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryWaitEventsResponse", string(data)}, " ")
}
