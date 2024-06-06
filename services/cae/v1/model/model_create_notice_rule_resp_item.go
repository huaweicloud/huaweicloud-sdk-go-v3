package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateNoticeRuleRespItem struct {

	// 通知规则的唯一标识。
	Id *string `json:"id,omitempty"`

	// 通知名称。
	Name *string `json:"name,omitempty"`

	// 触发事件名称。
	EventName *string `json:"event_name,omitempty"`

	Scope *NoticeRuleScope `json:"scope,omitempty"`

	TriggerPolicy *TriggerPolicy `json:"trigger_policy,omitempty"`

	Notification *NoticeRuleNotification `json:"notification,omitempty"`

	// 是否启用。
	Enable *bool `json:"enable,omitempty"`
}

func (o CreateNoticeRuleRespItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNoticeRuleRespItem struct{}"
	}

	return strings.Join([]string{"CreateNoticeRuleRespItem", string(data)}, " ")
}
