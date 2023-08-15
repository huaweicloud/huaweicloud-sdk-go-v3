package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AimSendTaskRequestMode 发送智能信息响应。
type AimSendTaskRequestMode struct {

	// 智能信息发送任务名称。
	TaskName *string `json:"task_name,omitempty"`

	SmsChannel *SmsChannel `json:"sms_channel,omitempty"`

	ResolveTask *AimResolveTaskRequestMode `json:"resolve_task,omitempty"`
}

func (o AimSendTaskRequestMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimSendTaskRequestMode struct{}"
	}

	return strings.Join([]string{"AimSendTaskRequestMode", string(data)}, " ")
}
