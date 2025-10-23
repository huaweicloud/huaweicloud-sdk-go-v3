package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventsResponse Response Object
type ListEventsResponse struct {

	// 本次分页查询响应的事件条数
	Count *int64 `json:"count,omitempty"`

	// 下一页的marker
	NextMarker *string `json:"next_marker,omitempty"`

	// 事件
	Events         *[]EventEntity `json:"events,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsResponse struct{}"
	}

	return strings.Join([]string{"ListEventsResponse", string(data)}, " ")
}
