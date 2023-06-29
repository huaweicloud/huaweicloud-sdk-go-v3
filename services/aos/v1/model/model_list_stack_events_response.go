package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStackEventsResponse Response Object
type ListStackEventsResponse struct {

	// 资源栈事件列表
	StackEvents    *[]StackEvent `json:"stack_events,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListStackEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStackEventsResponse struct{}"
	}

	return strings.Join([]string{"ListStackEventsResponse", string(data)}, " ")
}
