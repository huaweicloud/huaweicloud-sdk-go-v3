package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserEventsRequest Request Object
type ListUserEventsRequest struct {

	// 查询起始时间(0时区)
	StartTime *string `json:"start_time,omitempty"`

	// 查询结束时间(0时区)
	EndTime *string `json:"end_time,omitempty"`

	// 用户名（精确搜索）
	Username *string `json:"username,omitempty"`

	// 事件类型
	EventType *string `json:"event_type,omitempty"`

	// 事件之间的关联id
	EventTraceId *string `json:"event_trace_id,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，返回用户事件数量限制，取值范围0-1000。如果不指定，默认为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListUserEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserEventsRequest struct{}"
	}

	return strings.Join([]string{"ListUserEventsRequest", string(data)}, " ")
}
