package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmConfig struct {
	AccountName *string `json:"account_name,omitempty"`

	AlarmId *string `json:"alarm_id,omitempty"`

	// 告警周期，0：8时到22时，1：全天
	AlarmTimePeriod *int32 `json:"alarm_time_period,omitempty"`

	// 告警类型 0:攻击告警; 1:流量超额预警; 2:EIP未防护告警; 3:威胁情报告警
	AlarmType *int32 `json:"alarm_type,omitempty"`

	// 告警状态 0:失效; 1:生效
	EnableStatus *int32 `json:"enable_status,omitempty"`

	// 告警触发频次
	FrequencyCount *int32 `json:"frequency_count,omitempty"`

	// 告警频次时间范围，以分钟为单位
	FrequencyTime *int32 `json:"frequency_time,omitempty"`

	// 告警语言，zh-cn为中文，en-us为英文
	Language *string `json:"language,omitempty"`

	// 通知群组名称
	Name *string `json:"name,omitempty"`

	// 告警等级，当type为0和4时，severity为CRITICAL,HIGH,MEDIUM,LOW四种等级的组合字符串，当type为2时，severity固定为3
	Severity *string `json:"severity,omitempty"`

	// 告警topic的urn
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 用户名称，为cfw
	Username *string `json:"username,omitempty"`
}

func (o AlarmConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmConfig struct{}"
	}

	return strings.Join([]string{"AlarmConfig", string(data)}, " ")
}
