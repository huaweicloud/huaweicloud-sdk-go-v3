package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmTopicConfigInfoNewResponse Response Object
type ListAlarmTopicConfigInfoNewResponse struct {

	// CPU告警阈值(%)
	AlarmCpu *string `json:"alarm_cpu,omitempty"`

	// 磁盘告警阈值(%)
	AlarmDisk *string `json:"alarm_disk,omitempty"`

	// 内存告警阈值(%)
	AlarmMemory *string `json:"alarm_memory,omitempty"`

	// 每天发送告警总条数
	AlarmNum *string `json:"alarm_num,omitempty"`

	// 告警等级 - high：高  - medium：中  - low：低
	AlarmRisk *[]string `json:"alarm_risk,omitempty"`

	// 通知开关 - ON：开启 - OFF：关闭
	AlarmSwitch *string `json:"alarm_switch,omitempty"`

	// 通知消息主题URN,调用SMN服务接口获取
	AlarmTopicUrn *string `json:"alarm_topic_urn,omitempty"`

	// 是否支持SMN订阅  - true：支持  - false：不支持
	SmnEffective   *bool `json:"smn_effective,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListAlarmTopicConfigInfoNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmTopicConfigInfoNewResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmTopicConfigInfoNewResponse", string(data)}, " ")
}
