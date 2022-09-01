package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEventDetailResponse struct {
	EventInfo      *SpanEventInfo `json:"event_info,omitempty" xml:"event_info"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowEventDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowEventDetailResponse", string(data)}, " ")
}
