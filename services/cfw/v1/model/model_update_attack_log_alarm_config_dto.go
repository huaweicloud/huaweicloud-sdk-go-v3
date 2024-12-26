package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAttackLogAlarmConfigDto struct {

	// 账号名称
	AccountName *string `json:"account_name,omitempty"`

	// 告警id
	AlarmId *string `json:"alarm_id,omitempty"`

	// 告警周期，0：全天，1：8时到22时
	AlarmTimePeriod *int32 `json:"alarm_time_period,omitempty"`

	// 告警类型 0:攻击告警; 1:流量超额预警; 2:EIP未防护告警; 3:威胁情报告警
	AlarmType *int32 `json:"alarm_type,omitempty"`

	// 告警状态 0:失效; 1:生效
	EnableStatus *int32 `json:"enable_status,omitempty"`

	// 告警触发频次
	FrequencyCount *int32 `json:"frequency_count,omitempty"`

	// 告警频次时间范围
	FrequencyTime *int32 `json:"frequency_time,omitempty"`

	// 告警语言
	Language *string `json:"language,omitempty"`

	// 告警等级
	Severity *string `json:"severity,omitempty"`

	// 告警urn
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 用户名称
	Username *string `json:"username,omitempty"`
}

func (o UpdateAttackLogAlarmConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAttackLogAlarmConfigDto struct{}"
	}

	return strings.Join([]string{"UpdateAttackLogAlarmConfigDto", string(data)}, " ")
}
