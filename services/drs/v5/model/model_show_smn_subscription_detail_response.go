package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSmnSubscriptionDetailResponse Response Object
type ShowSmnSubscriptionDetailResponse struct {

	// SMN主题名称
	TopicName *string `json:"topic_name,omitempty"`

	// SMN主题的唯一资源标识
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 延迟时间，单位秒
	DelayTime *int32 `json:"delay_time,omitempty"`

	// RPO延迟时间，单位秒
	RpoDelay *int32 `json:"rpo_delay,omitempty"`

	// RTO延迟时间，单位秒
	RtoDelay *int32 `json:"rto_delay,omitempty"`

	// 是否发送告警给订阅用户
	IsAlarmUser    *bool `json:"is_alarm_user,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowSmnSubscriptionDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSmnSubscriptionDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowSmnSubscriptionDetailResponse", string(data)}, " ")
}
