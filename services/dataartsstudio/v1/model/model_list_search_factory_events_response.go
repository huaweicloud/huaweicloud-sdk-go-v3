package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSearchFactoryEventsResponse Response Object
type ListSearchFactoryEventsResponse struct {

	// 事件。
	Events *[]EventSearchResultV2Events `json:"events,omitempty"`

	// 总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSearchFactoryEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSearchFactoryEventsResponse struct{}"
	}

	return strings.Join([]string{"ListSearchFactoryEventsResponse", string(data)}, " ")
}
