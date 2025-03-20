package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PersonnelResponse 人员信息
type PersonnelResponse struct {

	// 该用户是否是根用户
	IsRootUser *bool `json:"is_root_user,omitempty"`

	// 钉钉回调
	DingtalkWebhook *string `json:"dingtalk_webhook,omitempty"`

	// 邮箱
	Email *string `json:"email,omitempty"`

	// 用户id
	Id *string `json:"id,omitempty"`

	// 手机
	Mobile *string `json:"mobile,omitempty"`

	// 用户名
	Name *string `json:"name,omitempty"`

	// 企业微信回调
	WecomWebhook *string `json:"wecom_webhook,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 钉钉秘钥
	DingtalkPrivateKey *string `json:"dingtalk_private_key,omitempty"`

	// 短信订阅状态
	MsgSubscriptionStatus *int32 `json:"msg_subscription_status,omitempty"`

	// 企业微信订阅状态
	WeichatSubscriptionStatus *int32 `json:"weichat_subscription_status,omitempty"`

	// 钉钉订阅状态
	DingTalkSubscriptionStatus *int32 `json:"ding_talk_subscription_status,omitempty"`

	// 邮箱订阅
	EmailSubscriptionStatus *int32 `json:"email_subscription_status,omitempty"`

	// 语音订阅状态
	CallNotifySubscriptionStatus *int32 `json:"call_notify_subscription_status,omitempty"`

	// 账号来源类型
	SourceType *string `json:"source_type,omitempty"`

	// 账号状态
	SourceStatus *string `json:"source_status,omitempty"`

	// 飞书回调
	LarkWebhook *string `json:"lark_webhook,omitempty"`

	// 飞书订阅状态
	LarkSubscriptionStatus *int32 `json:"lark_subscription_status,omitempty"`
}

func (o PersonnelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersonnelResponse struct{}"
	}

	return strings.Join([]string{"PersonnelResponse", string(data)}, " ")
}
