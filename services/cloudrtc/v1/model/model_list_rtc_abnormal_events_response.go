package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRtcAbnormalEventsResponse struct {
	// 异常总数

	Total *int32 `json:"total,omitempty"`
	// 异常体验列表

	Events *[]AbnormalEvent `json:"events,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcAbnormalEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcAbnormalEventsResponse struct{}"
	}

	return strings.Join([]string{"ListRtcAbnormalEventsResponse", string(data)}, " ")
}
