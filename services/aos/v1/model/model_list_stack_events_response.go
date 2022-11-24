package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListStackEventsResponse struct {

	// 栈的更新状态
	StackEvents *[]StackEventResponse `json:"stack_events,omitempty"`

	// 当一页无法发回所有的细节，将返回next_marker，客户可以继续调用list-stack-events并给与next_marker来继续读取下页
	NextMarker     *string `json:"next_marker,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListStackEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStackEventsResponse struct{}"
	}

	return strings.Join([]string{"ListStackEventsResponse", string(data)}, " ")
}
