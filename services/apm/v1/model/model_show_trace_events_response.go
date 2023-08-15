package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTraceEventsResponse Response Object
type ShowTraceEventsResponse struct {

	// span event信息列表。
	SpanEventList  *[]SpanEventInfo `json:"span_event_list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowTraceEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTraceEventsResponse struct{}"
	}

	return strings.Join([]string{"ShowTraceEventsResponse", string(data)}, " ")
}
