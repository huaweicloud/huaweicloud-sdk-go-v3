package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledEventsRequest Request Object
type ListScheduledEventsRequest struct {

	// 每页返回的事件个数。取值范围：1~100（默认值为10）
	Limit *int32 `json:"limit,omitempty"`

	// 分页标记。从marker指定的下一条数据开始查询。
	Marker *string `json:"marker,omitempty"`

	// 计划事件id。
	EventId *string `json:"event_id,omitempty"`

	// 云手机服务器的唯一标识。
	ServerId *string `json:"server_id,omitempty"`

	// 事件发布开始时间，按照时间范围过滤。
	PublishSince *string `json:"publish_since,omitempty"`

	// 事件发布结束时间，按照时间范围过滤。
	PublishUntil *string `json:"publish_until,omitempty"`

	// 计划事件状态。支持多值查询过滤。 取值范围： inquiring: 待授权、 scheduled：待执行、 executing：执行中、 completed：执行成功、 failed：执行失败、 canceled：取消
	State *[]string `json:"state,omitempty"`

	// 计划事件类型。支持多值查询过滤。取值范围： localdisk-recovery：本地盘换盘、 system-maintenance：系统维护
	Type *[]string `json:"type,omitempty"`
}

func (o ListScheduledEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledEventsRequest struct{}"
	}

	return strings.Join([]string{"ListScheduledEventsRequest", string(data)}, " ")
}
