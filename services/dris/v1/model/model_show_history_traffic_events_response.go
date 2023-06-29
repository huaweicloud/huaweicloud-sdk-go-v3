package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHistoryTrafficEventsResponse Response Object
type ShowHistoryTrafficEventsResponse struct {

	// **参数说明**：条件查询返回的总条数。
	Count *int64 `json:"count,omitempty"`

	// **参数说明**：事件列表。
	Events         *[]HistoryTrafficEventDto `json:"events,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowHistoryTrafficEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHistoryTrafficEventsResponse struct{}"
	}

	return strings.Join([]string{"ShowHistoryTrafficEventsResponse", string(data)}, " ")
}
