package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventStreamingResponse Response Object
type ListEventStreamingResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本页数量
	Size *int32 `json:"size,omitempty"`

	// 对象列表
	Items *[]EventStreamingDetail `json:"items,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListEventStreamingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventStreamingResponse struct{}"
	}

	return strings.Join([]string{"ListEventStreamingResponse", string(data)}, " ")
}
