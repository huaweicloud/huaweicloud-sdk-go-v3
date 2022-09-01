package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRtcAbnormalEventsResponse struct {

	// 异常总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 异常体验列表
	Events *[]AbnormalEvent `json:"events,omitempty" xml:"events"`

	XRequestId     *string `json:"X-request-id,omitempty" xml:"X-request-id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcAbnormalEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcAbnormalEventsResponse struct{}"
	}

	return strings.Join([]string{"ListRtcAbnormalEventsResponse", string(data)}, " ")
}
