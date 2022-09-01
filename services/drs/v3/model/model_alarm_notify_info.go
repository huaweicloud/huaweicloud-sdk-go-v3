package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 异常通知信息
type AlarmNotifyInfo struct {

	// 订阅延迟时间(单位为s)
	DelayTime *int64 `json:"delay_time,omitempty" xml:"delay_time"`

	// rto延迟时间
	RtoDelay *int64 `json:"rto_delay,omitempty" xml:"rto_delay"`

	// rpo延迟时间
	RpoDelay *int64 `json:"rpo_delay,omitempty" xml:"rpo_delay"`

	// 异常告警是否通知用户
	AlarmToUser *bool `json:"alarm_to_user,omitempty" xml:"alarm_to_user"`

	// 收件方式与信息体
	Subscriptions *[]SubscriptionInfo `json:"subscriptions,omitempty" xml:"subscriptions"`
}

func (o AlarmNotifyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmNotifyInfo struct{}"
	}

	return strings.Join([]string{"AlarmNotifyInfo", string(data)}, " ")
}
