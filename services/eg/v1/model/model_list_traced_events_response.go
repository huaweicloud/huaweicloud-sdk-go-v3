package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTracedEventsResponse Response Object
type ListTracedEventsResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本页数量
	Size *int32 `json:"size,omitempty"`

	// 事件追踪列表
	Result *[]ListTracedEventsRespResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTracedEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTracedEventsResponse struct{}"
	}

	return strings.Join([]string{"ListTracedEventsResponse", string(data)}, " ")
}
