package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledEventsRequest Request Object
type ListScheduledEventsRequest struct {

	// 从marker指定的事件的下一条数据开始查询。
	Marker *string `json:"marker,omitempty"`

	// 事件ID
	Id *string `json:"id,omitempty"`

	// 实例ID
	InstanceId *[]string `json:"instance_id,omitempty"`

	// 事件类型
	Type *[]string `json:"type,omitempty"`

	// 事件状态
	State *[]string `json:"state,omitempty"`

	// 事件发布开始时间
	PublishSince *string `json:"publish_since,omitempty"`

	// 事件发布截至时间
	PublishUntil *string `json:"publish_until,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListScheduledEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledEventsRequest struct{}"
	}

	return strings.Join([]string{"ListScheduledEventsRequest", string(data)}, " ")
}
