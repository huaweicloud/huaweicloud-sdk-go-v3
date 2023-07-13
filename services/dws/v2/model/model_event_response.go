package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventResponse 事件返回体
type EventResponse struct {

	// 事件类别
	Category *string `json:"category,omitempty"`

	// 事件描述
	Description *string `json:"description,omitempty"`

	// 事件ID
	EventId *string `json:"event_id,omitempty"`

	// 事件定义名称
	Name *string `json:"name,omitempty"`

	// 事件显示名称
	DisplayName *string `json:"display_name,omitempty"`

	// 所属服务
	NameSpace *string `json:"name_space,omitempty"`

	// 事件级别
	Severity *string `json:"severity,omitempty"`

	// 事件源类别
	SourceType *string `json:"source_type,omitempty"`

	// 时间
	OccurTime *int64 `json:"occur_time,omitempty"`

	// 租户凭证ID
	ProjectId *string `json:"project_id,omitempty"`

	// 事件源ID
	SourceId *string `json:"source_id,omitempty"`

	// 事件源名称
	SourceName *string `json:"source_name,omitempty"`

	// 状态
	Status *int32 `json:"status,omitempty"`

	// 事件主题
	Subject *string `json:"subject,omitempty"`

	// 事件信息
	Context *string `json:"context,omitempty"`
}

func (o EventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventResponse struct{}"
	}

	return strings.Join([]string{"EventResponse", string(data)}, " ")
}
