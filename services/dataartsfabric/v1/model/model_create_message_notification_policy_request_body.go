package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMessageNotificationPolicyRequestBody 创建消息通知策略请求体
type CreateMessageNotificationPolicyRequestBody struct {

	// 消息通知策略
	MessageNotificationPolicyList []CreateMessageNotificationPolicy `json:"message_notification_policy_list"`
}

func (o CreateMessageNotificationPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageNotificationPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateMessageNotificationPolicyRequestBody", string(data)}, " ")
}
