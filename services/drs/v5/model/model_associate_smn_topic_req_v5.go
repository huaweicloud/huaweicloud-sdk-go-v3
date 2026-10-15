package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateSmnTopicReqV5 关联SMN主题请求体
type AssociateSmnTopicReqV5 struct {

	// 主题资源标识
	TopicUrn string `json:"topic_urn"`

	// 增量时延阈值, 单位：秒（s）
	IncrementDelayThreshold *int64 `json:"increment_delay_threshold,omitempty"`

	// RTO时延阈值，灾备场景下可配置，单位：秒（s）
	RtoDelayThreshold *int64 `json:"rto_delay_threshold,omitempty"`

	// RPO时延阈值，灾备场景下可配置，单位：秒（s）
	RpoDelayThreshold *int64 `json:"rpo_delay_threshold,omitempty"`

	// 异常告警是否通知用户，取值范围[true,false]
	IsAlarmToUser bool `json:"is_alarm_to_user"`
}

func (o AssociateSmnTopicReqV5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSmnTopicReqV5 struct{}"
	}

	return strings.Join([]string{"AssociateSmnTopicReqV5", string(data)}, " ")
}
