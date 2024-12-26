package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventSampleItemInfo struct {

	// 事件示例ID
	Id *string `json:"id,omitempty"`

	// 事件示例名称
	Name *string `json:"name,omitempty"`

	// 事件示例内容
	Content *string `json:"content,omitempty"`

	// 事件示例对应的事件类型ID
	EventTypeId *string `json:"event_type_id,omitempty"`

	// 事件示例对应的事件类型名称
	EventTypeName *string `json:"event_type_name,omitempty"`

	// 事件示例对应的事件源ID
	EventSourceId *string `json:"event_source_id,omitempty"`

	// 事件示例对应的事件源名称
	EventSourceName *string `json:"event_source_name,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 删除时间
	DeletedTime *string `json:"deleted_time,omitempty"`
}

func (o EventSampleItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventSampleItemInfo struct{}"
	}

	return strings.Join([]string{"EventSampleItemInfo", string(data)}, " ")
}
