package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AimSendTaskInfo 发送任务消息体。
type AimSendTaskInfo struct {

	// 智能信息发送任务名称。
	TaskName *string `json:"task_name,omitempty"`

	SmsChannel *AimSendTaskSmsChannel `json:"sms_channel,omitempty"`

	ResolveTask *AimResolveTaskMode `json:"resolve_task,omitempty"`
}

func (o AimSendTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AimSendTaskInfo struct{}"
	}

	return strings.Join([]string{"AimSendTaskInfo", string(data)}, " ")
}
