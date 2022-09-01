package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改收件方式与信息体
type BatchSetAlarmNotifyInfo struct {

	// 手动输入手机号、邮箱模式时填写
	Subscriptions *[]SubscriptionInfo `json:"subscriptions,omitempty" xml:"subscriptions"`

	// 主题资源标识
	TopicUrn *string `json:"topic_urn,omitempty" xml:"topic_urn"`

	// 订阅延迟时间
	DelayTime *int64 `json:"delay_time,omitempty" xml:"delay_time"`

	// rto延迟时间
	RtoDelay *int64 `json:"rto_delay,omitempty" xml:"rto_delay"`

	// rpo延迟时间
	RpoDelay *int64 `json:"rpo_delay,omitempty" xml:"rpo_delay"`

	// 异常告警是否通知用户，不填默认为false
	AlarmToUser *bool `json:"alarm_to_user,omitempty" xml:"alarm_to_user"`
}

func (o BatchSetAlarmNotifyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetAlarmNotifyInfo struct{}"
	}

	return strings.Join([]string{"BatchSetAlarmNotifyInfo", string(data)}, " ")
}
