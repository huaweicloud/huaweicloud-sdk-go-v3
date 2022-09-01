package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEventsResponse struct {

	// 测试事件总数。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 测试事件列表。
	Events *[]ListEventsResult `json:"events,omitempty" xml:"events"`

	// 下次读取位置。
	NextMarker     *int64 `json:"next_marker,omitempty" xml:"next_marker"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventsResponse struct{}"
	}

	return strings.Join([]string{"ListEventsResponse", string(data)}, " ")
}
