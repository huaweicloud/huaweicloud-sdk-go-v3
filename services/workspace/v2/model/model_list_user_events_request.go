package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserEventsRequest Request Object
type ListUserEventsRequest struct {

	// 查询起始时间，格式为：UTC时间，例如\"1970-01-01T00:00:00Z\"。
	StartTime string `json:"start_time"`

	// 查询结束时间，格式为：UTC时间，例如\"1970-01-01T00:00:00Z\"。
	EndTime string `json:"end_time"`

	// 用户名（精确搜索）。
	Username *string `json:"username,omitempty"`

	// 事件类型，英文逗号隔开。
	EventType *string `json:"event_type,omitempty"`

	// 操作资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 操作资源名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
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
