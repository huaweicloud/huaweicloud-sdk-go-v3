package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchShowTrafficEventsResponse struct {

	// **参数说明**：返回条目总数量。
	Count *int64 `json:"count,omitempty"`

	// **参数说明**：事件列表。
	Events         *[]TrafficEventDto `json:"events,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o BatchShowTrafficEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowTrafficEventsResponse struct{}"
	}

	return strings.Join([]string{"BatchShowTrafficEventsResponse", string(data)}, " ")
}
