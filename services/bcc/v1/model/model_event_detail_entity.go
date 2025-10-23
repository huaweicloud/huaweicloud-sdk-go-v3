package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventDetailEntity 事件详情实体类
type EventDetailEntity struct {

	// 事件内容
	Content *string `json:"content,omitempty"`

	// 组ID
	GroupId *string `json:"group_id,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名
	ResourceName *string `json:"resource_name,omitempty"`

	EventState *EventStateEnum `json:"event_state,omitempty"`

	EventLevel *EventLevelEnum `json:"event_level,omitempty"`

	// 事件用户
	EventUser *string `json:"event_user,omitempty"`

	// 事件类型
	EventType *string `json:"event_type,omitempty"`

	// 事件维度
	Dimensions *string `json:"dimensions,omitempty"`
}

func (o EventDetailEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventDetailEntity struct{}"
	}

	return strings.Join([]string{"EventDetailEntity", string(data)}, " ")
}
