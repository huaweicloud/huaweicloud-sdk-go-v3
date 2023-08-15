package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetAlarmNotifyInfo 修改收件方式与信息体
type BatchSetAlarmNotifyInfo struct {

	// 手动输入手机号、邮箱模式时填写
	Subscriptions *[]SubscriptionInfo `json:"subscriptions,omitempty"`

	// 主题资源标识
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 订阅延迟时间
	DelayTime *int64 `json:"delay_time,omitempty"`

	// rto延迟时间
	RtoDelay *int64 `json:"rto_delay,omitempty"`

	// rpo延迟时间
	RpoDelay *int64 `json:"rpo_delay,omitempty"`

	// 异常告警是否通知用户，不填默认为false
	AlarmToUser *bool `json:"alarm_to_user,omitempty"`
}

func (o BatchSetAlarmNotifyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetAlarmNotifyInfo struct{}"
	}

	return strings.Join([]string{"BatchSetAlarmNotifyInfo", string(data)}, " ")
}
