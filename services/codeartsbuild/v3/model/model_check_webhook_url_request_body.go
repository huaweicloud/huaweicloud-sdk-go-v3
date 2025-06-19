package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckWebhookUrlRequestBody 检查webhook地址
type CheckWebhookUrlRequestBody struct {

	// 任务ID
	JobId string `json:"job_id"`

	// 通知的类型,分为消息，邮件，钉钉，飞书和微信
	NoticeType string `json:"notice_type"`

	// Webhook地址
	WebhookUrl string `json:"webhook_url"`
}

func (o CheckWebhookUrlRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWebhookUrlRequestBody struct{}"
	}

	return strings.Join([]string{"CheckWebhookUrlRequestBody", string(data)}, " ")
}
