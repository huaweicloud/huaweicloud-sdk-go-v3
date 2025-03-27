package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EventResponse struct {

	// 事件ID
	Id string `json:"id"`

	// 事件类型
	Type *string `json:"type,omitempty"`

	// 事件状态
	State *string `json:"state,omitempty"`

	// 事件发布时间
	PublishTime *string `json:"publish_time,omitempty"`

	// 事件开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 事件完成时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 事件计划执行开始时间
	NotBefore *string `json:"not_before,omitempty"`

	// 事件计划执行完成时间
	NotAfter *string `json:"not_after,omitempty"`

	// 事件计划执行开始时间deadline
	NotBeforeDeadline *string `json:"not_before_deadline,omitempty"`

	// 事件描述
	Description *string `json:"description,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	ExecuteOptions *EventResponseExecuteOptions `json:"execute_options,omitempty"`

	Source *EventResponseSource `json:"source,omitempty"`
}

func (o EventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventResponse struct{}"
	}

	return strings.Join([]string{"EventResponse", string(data)}, " ")
}
