package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserEventsResponse Response Object
type ListUserEventsResponse struct {

	// 事件总数。
	Count *int32 `json:"count,omitempty"`

	// 用户事件列表。
	Items          *[]UserEventRsp `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListUserEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserEventsResponse struct{}"
	}

	return strings.Join([]string{"ListUserEventsResponse", string(data)}, " ")
}
