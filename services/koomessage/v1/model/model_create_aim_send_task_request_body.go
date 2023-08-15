package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAimSendTaskRequestBody 发送AIM消息体。
type CreateAimSendTaskRequestBody struct {

	// 智能信息发送任务名称。  > 不能为空白字符串。
	TaskName string `json:"task_name"`

	SmsChannel *SmsChannel `json:"sms_channel"`

	ResolveTask *AimResolveTask `json:"resolve_task"`
}

func (o CreateAimSendTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAimSendTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAimSendTaskRequestBody", string(data)}, " ")
}
