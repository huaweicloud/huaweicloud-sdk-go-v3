package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRtcEventResponse Response Object
type ListRtcEventResponse struct {

	// 时间
	Ctime *string `json:"ctime,omitempty"`

	// 异常事件ID
	EventId *string `json:"event_id,omitempty"`

	// 异常事件信息
	EventInfo *string `json:"event_info,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcEventResponse struct{}"
	}

	return strings.Join([]string{"ListRtcEventResponse", string(data)}, " ")
}
